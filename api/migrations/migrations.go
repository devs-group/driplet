package migrations

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pressly/goose/v3"
)

//go:embed *.sql
var embedMigrations embed.FS

// RunMigrations runs the migrations based on the command
func RunMigrations(db *sql.DB, command string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect: %v", err)
	}

	goose.SetBaseFS(embedMigrations)

	switch command {
	case "up":
		if err := goose.Up(db, "."); err != nil {
			return fmt.Errorf("failed to run migrations: %v", err)
		}
	case "down":
		if err := goose.Down(db, "."); err != nil {
			return fmt.Errorf("failed to rollback migration: %v", err)
		}
	case "reset":
		if err := goose.Reset(db, "."); err != nil {
			return fmt.Errorf("failed to reset migrations: %v", err)
		}
	case "status":
		if err := goose.Status(db, "."); err != nil {
			return fmt.Errorf("failed to get migrations status: %v", err)
		}
	default:
		return fmt.Errorf("unknown command: %s", command)
	}
	return nil
}

// CreateMigrationFile creates a new migration file in the migrations directory
func CreateMigrationFile(name string) error {
	timestamp := time.Now().UTC().Format("20060102150405")
	filename := fmt.Sprintf("%s_%s.sql", timestamp, name)

	// Get the absolute path to the migrations directory
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	migrationsDir := filepath.Join(dir, "migrations")
	if err := os.MkdirAll(migrationsDir, 0755); err != nil {
		return fmt.Errorf("failed to create migrations directory: %w", err)
	}

	filepath := filepath.Join(migrationsDir, filename)

	content := fmt.Sprintf(`-- +goose Up
-- +goose StatementBegin
SELECT 1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 1;
-- +goose StatementEnd
`)

	if err := os.WriteFile(filepath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to create migration file: %w", err)
	}

	fmt.Printf("Created new migration: %s\n", filename)
	return nil
}
