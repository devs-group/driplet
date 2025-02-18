package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Config struct {
	ConnectionString string
	MaxOpenConns     int
	MaxIdleConns     int
	ConnMaxLifetime  time.Duration
	ConnMaxIdleTime  time.Duration
	// Retry configuration
	MaxRetries    int
	RetryDelay    time.Duration
	RetryMaxDelay time.Duration
}

// DefaultConfig returns a default database configuration based on environment variables
func DefaultConfig() Config {
	// Get connection parameters from environment variables with defaults
	user := getEnv("POSTGRES_USER", "postgres")
	password := getEnv("POSTGRES_PASSWORD", "postgres")
	dbName := getEnv("POSTGRES_DB", "postgres")
	host := getEnv("POSTGRES_HOST", "database")
	port := getEnv("POSTGRES_PORT", "5432")
	sslMode := getEnv("POSTGRES_SSLMODE", "disable")

	// Build connection string
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbName, sslMode,
	)

	// Parse connection settings
	maxOpenConns := getIntEnv("DB_MAX_OPEN_CONNS", 10)
	maxIdleConns := getIntEnv("DB_MAX_IDLE_CONNS", 5)
	connMaxLifetime := getDurationEnv("DB_CONN_MAX_LIFETIME", time.Hour)
	connMaxIdleTime := getDurationEnv("DB_CONN_MAX_IDLE_TIME", 15*time.Minute)

	// Get retry configuration
	maxRetries := getIntEnv("DB_MAX_RETRIES", 5)
	retryDelay := getDurationEnv("DB_RETRY_DELAY", 1*time.Second)
	retryMaxDelay := getDurationEnv("DB_RETRY_MAX_DELAY", 30*time.Second)

	return Config{
		ConnectionString: connString,
		MaxOpenConns:     maxOpenConns,
		MaxIdleConns:     maxIdleConns,
		ConnMaxLifetime:  connMaxLifetime,
		ConnMaxIdleTime:  connMaxIdleTime,
		MaxRetries:       maxRetries,
		RetryDelay:       retryDelay,
		RetryMaxDelay:    retryMaxDelay,
	}
}

type Database struct {
	SQLX *sqlx.DB
	SQL  *sql.DB
}

func Connect(config Config) (*Database, error) {
	slog.Info("connecting to postgres database",
		"host", getHostFromConnString(config.ConnectionString),
		"max_conns", config.MaxOpenConns,
		"max_retries", config.MaxRetries)

	var db *sqlx.DB
	var err error
	var currentDelay time.Duration = config.RetryDelay

	// Retry logic for establishing the connection
	for attempt := 0; attempt <= config.MaxRetries; attempt++ {
		if attempt > 0 {
			slog.Info("retrying database connection",
				"attempt", attempt,
				"max_attempts", config.MaxRetries,
				"delay", currentDelay)

			time.Sleep(currentDelay)
			// Exponential backoff with cap
			currentDelay *= 2
			if currentDelay > config.RetryMaxDelay {
				currentDelay = config.RetryMaxDelay
			}
		}

		db, err = sqlx.Connect("postgres", config.ConnectionString)
		if err == nil {
			break
		}

		slog.Warn("database connection failed",
			"attempt", attempt,
			"error", err.Error())
		if attempt == config.MaxRetries {
			return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", config.MaxRetries+1, err)
		}
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)

	slog.Info("verifying database connection with ping")

	// Verify connection with retries
	currentDelay = config.RetryDelay
	for attempt := 0; attempt <= config.MaxRetries; attempt++ {
		if attempt > 0 {
			slog.Info("retrying database ping",
				"attempt", attempt,
				"max_attempts", config.MaxRetries,
				"delay", currentDelay)

			time.Sleep(currentDelay)
			// Exponential backoff with cap
			currentDelay *= 2
			if currentDelay > config.RetryMaxDelay {
				currentDelay = config.RetryMaxDelay
			}
		}

		err = db.Ping()
		if err == nil {
			break
		}

		slog.Warn("database ping failed",
			"attempt", attempt,
			"error", err.Error())

		if attempt == config.MaxRetries {
			db.Close() // Close the connection before returning error
			return nil, fmt.Errorf("failed to ping database after %d attempts: %w", config.MaxRetries+1, err)
		}
	}

	slog.Info("database connection has been established successfully!")

	return &Database{
		SQLX: db,
		SQL:  db.DB,
	}, nil
}

// Close closes the database connection
func (d *Database) Close() error {
	slog.Info("closing database connection")
	return d.SQLX.Close()
}

// Helper functions for environment variables and connection

// getHostFromConnString extracts the host from a connection string
func getHostFromConnString(connString string) string {
	// This is a simple implementation - it looks for "host=" in the connection string
	parts := strings.Split(connString, " ")
	for _, part := range parts {
		if strings.HasPrefix(part, "host=") {
			return strings.TrimPrefix(part, "host=")
		}
	}
	return "unknown"
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

// getIntEnv retrieves an environment variable as an integer or returns a default value
func getIntEnv(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// getDurationEnv retrieves an environment variable as a duration or returns a default value
func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := time.ParseDuration(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
