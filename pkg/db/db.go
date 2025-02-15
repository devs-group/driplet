package db

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
	SSLMode      string
	MaxRetries   int
	RetryDelay   time.Duration
}

func DefaultConfig() Config {
	return Config{
		Host:         getEnvOrDefault("DB_HOST", "localhost"),
		Port:         getEnvOrDefault("DB_PORT", "5432"),
		User:         getEnvOrDefault("DB_USER", "postgres"),
		Password:     getEnvOrDefault("DB_PASSWORD", "postgres"),
		DatabaseName: getEnvOrDefault("DB_NAME", "postgres"),
		SSLMode:      getEnvOrDefault("DB_SSLMODE", "disable"),
		MaxRetries:   5,
		RetryDelay:   time.Second * 5,
	}
}

// Connect establishes a connection to the database with retries
func Connect(cfg Config) (*sqlx.DB, error) {
	slog.Info("connecting to the db...", "host", cfg.Host, "port", cfg.Port, "user", cfg.User, "db", cfg.DatabaseName, "sslMode", cfg.SSLMode)
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DatabaseName, cfg.SSLMode,
	)

	var db *sqlx.DB
	var err error

	for i := 0; i < cfg.MaxRetries; i++ {
		db, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			break
		}

		if i < cfg.MaxRetries-1 {
			time.Sleep(cfg.RetryDelay)
			continue
		}
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", cfg.MaxRetries, err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	slog.Info("connecion to the db has been established!")

	return db, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
