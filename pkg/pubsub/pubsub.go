package pubsub

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type Client struct {
	*pubsub.Client
	projectID string
}

type Config struct {
	ProjectID               string
	CredentialsFile         string
	EmulatorHost            string
	AutoCreateTopics        bool
	DefaultSubscriberConfig SubscriberConfig
}

type SubscriberConfig struct {
	MaxOutstandingMessages int
	NumGoroutines          int
	AsyncPull              bool
}

func DefaultConfig() Config {
	return Config{
		ProjectID:        getEnvOrDefault("PUBSUB_PROJECT_ID", "local-project"),
		CredentialsFile:  getEnvOrDefault("PUBSUB_APPLICATION_CREDENTIALS", ""),
		EmulatorHost:     os.Getenv("PUBSUB_EMULATOR_HOST"),
		AutoCreateTopics: true,
		DefaultSubscriberConfig: SubscriberConfig{
			MaxOutstandingMessages: 10,
			NumGoroutines:          1,
			AsyncPull:              false,
		},
	}
}

func Connect(ctx context.Context, cfg Config) (*Client, error) {
	var opts []option.ClientOption

	slog.Info("connecting to pubsub", "projectID", cfg.ProjectID)
	if cfg.EmulatorHost != "" {
		slog.Info("using pubsub emulator", "host", cfg.EmulatorHost)
		// Using emulator
		if err := os.Setenv("PUBSUB_EMULATOR_HOST", cfg.EmulatorHost); err != nil {
			return nil, fmt.Errorf("failed to set emulator host: %w", err)
		}
	} else if cfg.CredentialsFile != "" {
		slog.Info("using credentials file for pubsub", "path", cfg.CredentialsFile)
		// Using credentials file
		opts = append(opts, option.WithCredentialsFile(cfg.CredentialsFile))
	}

	client, err := pubsub.NewClient(ctx, cfg.ProjectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %w", err)
	}

	slog.Info("pubsub connection has been established!")
	return &Client{
		Client:    client,
		projectID: cfg.ProjectID,
	}, nil
}

type Publisher struct {
	topic *pubsub.Topic
	mu    sync.Mutex
}

func (c *Client) NewPublisher(topicID string, autoCreate bool) (*Publisher, error) {
	topic := c.Topic(topicID)

	if autoCreate {
		exists, err := topic.Exists(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to check if topic exists: %w", err)
		}

		if !exists {
			topic, err = c.CreateTopic(context.Background(), topicID)
			if err != nil {
				return nil, fmt.Errorf("failed to create topic: %w", err)
			}
			log.Printf("Created topic: %s", topicID)
		}
	}

	return &Publisher{
		topic: topic,
	}, nil
}

func (p *Publisher) Publish(ctx context.Context, data []byte, attrs map[string]string) (serverID string, err error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	msg := &pubsub.Message{
		Data:       data,
		Attributes: attrs,
	}

	result := p.topic.Publish(ctx, msg)
	serverID, err = result.Get(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to publish message: %w", err)
	}

	return serverID, nil
}

func (p *Publisher) Close() {
	p.topic.Stop()
}

type Subscriber struct {
	sub       *pubsub.Subscription
	config    SubscriberConfig
	projectID string
}

type MessageHandler func(context.Context, *pubsub.Message)

func (c *Client) NewSubscriber(topicID, subscriptionID string, cfg SubscriberConfig, autoCreate bool) (*Subscriber, error) {
	sub := c.Subscription(subscriptionID)

	if autoCreate {
		exists, err := sub.Exists(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to check if subscription exists: %w", err)
		}

		if !exists {
			topic := c.Topic(topicID)
			sub, err = c.CreateSubscription(context.Background(), subscriptionID, pubsub.SubscriptionConfig{
				Topic: topic,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to create subscription: %w", err)
			}
			log.Printf("Created subscription: %s", subscriptionID)
		}
	}

	return &Subscriber{
		sub:       sub,
		config:    cfg,
		projectID: c.projectID,
	}, nil
}

func (s *Subscriber) Subscribe(ctx context.Context, handler MessageHandler) error {
	// Configure subscription
	s.sub.ReceiveSettings.MaxOutstandingMessages = s.config.MaxOutstandingMessages
	s.sub.ReceiveSettings.NumGoroutines = s.config.NumGoroutines

	if s.config.AsyncPull {
		return s.sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
			go func() {
				handler(ctx, msg)
				msg.Ack()
			}()
		})
	}

	return s.sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		handler(ctx, msg)
		msg.Ack()
	})
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
