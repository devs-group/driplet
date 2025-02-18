package handlers

import (
	"context"
	"log/slog"

	"github.com/devs-group/driplet/api/di"
	"github.com/devs-group/driplet/pkg/pubsub"
	"github.com/devs-group/godi"
	"github.com/go-faster/errors"
	"github.com/gofiber/fiber/v2"
)

type EventsHandler struct {
	pubsubClient *pubsub.Client
}

func NewEventsHandler() (*EventsHandler, error) {
	pubsubClient, err := godi.Resolve[*pubsub.Client](di.Container)
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve pubsub client")
	}
	return &EventsHandler{pubsubClient: pubsubClient}, nil
}

func (h *EventsHandler) POST_CreateEvent(c *fiber.Ctx) error {
	publisher, err := h.pubsubClient.NewPublisher("client-events", true)
	if err != nil {
		slog.Error("failed to create publisher", "err", err)
		return fiber.ErrInternalServerError
	}
	ctx := context.Background()
	serverID, err := publisher.Publish(ctx, c.Body(), nil)
	if err != nil {
		slog.Error("failed to publish event", "err", err)
		return fiber.ErrInternalServerError
	}
	slog.Info("event has been published", "server_id", serverID)
	return c.JSON(fiber.Map{
		"server_id": serverID,
	})
}
