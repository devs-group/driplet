package calculate_points

import (
	"context"
	"fmt"
	"testing"
	"time"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculatePoints(t *testing.T) {
	testDBPort := uint32(5555)
	postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
		Port(testDBPort).
		Username("postgres").
		Password("postgres").
		Database("testdb"))

	err := postgres.Start()
	require.NoError(t, err, "Failed to start embedded Postgres")
	defer func() {
		err := postgres.Stop()
		require.NoError(t, err, "Failed to stop embedded Postgres")
	}()

	connStr := fmt.Sprintf("postgres://postgres:postgres@localhost:%d/testdb?sslmode=disable", testDBPort)
	db, err := sqlx.Connect("postgres", connStr)
	require.NoError(t, err, "Unable to connect to database")
	defer db.Close()

	// create test table
	_, err = db.Exec(`
		DROP TABLE IF EXISTS website_visits;
		CREATE TABLE website_visits (
			id SERIAL PRIMARY KEY,
			user_id VARCHAR(36) NOT NULL,
			timestamp TIMESTAMP NOT NULL,
			host VARCHAR(255) NOT NULL,
			path VARCHAR(255) NOT NULL
		);
	`)
	require.NoError(t, err, "Failed to create test table")

	// setup test data
	now := time.Now()
	yesterday := now.Add(-25 * time.Hour)

	testData := []struct {
		userID    string
		timestamp time.Time
		host      string
		path      string
	}{
		{"1", now.Add(-1 * time.Hour), "example.com", "/"},      // User 1: 1 point
		{"1", now.Add(-2 * time.Hour), "example.com", "/"},      // Duplicate - no additional points
		{"1", now.Add(-3 * time.Hour), "example.com", "/about"}, // Different path - 1 point

		{"2", now.Add(-1 * time.Hour), "example.com", "/"},  // User 2: 1 point
		{"2", now.Add(-2 * time.Hour), "example2.com", "/"}, // Different host - 1 point
		{"2", now.Add(-3 * time.Hour), "example3.com", "/"}, // Different host - 1 point

		{"3", yesterday, "example.com", "/"},                       // Outside 24 hours - no points
		{"3", now.Add(-23 * time.Hour), "example.com", "/contact"}, // Inside 24 hours - 1 point

		{"4", now.Add(-12 * time.Hour), "example.com", "/"},     // User 4: 1 point
		{"4", now.Add(-13 * time.Hour), "example.com", "/blog"}, // Different path - 1 point
		{"4", now.Add(-14 * time.Hour), "example.com", "/blog"}, // Duplicate - no additional points
	}

	for _, td := range testData {
		_, err = db.Exec(`
			INSERT INTO website_visits (user_id, timestamp, host, path)
			VALUES ($1, $2, $3, $4)
		`, td.userID, td.timestamp, td.host, td.path)
		require.NoError(t, err, "Failed to insert test data")
	}

	ctx := context.Background()
	userPoints, err := CalculatePoints(ctx, db)
	require.NoError(t, err, "CalculatePoints failed")

	// Expected results:
	// User 1: 2 points (example.com/ and example.com/about)
	// User 2: 3 points (example.com/, example2.com/, example3.com/)
	// User 3: 1 point (example.com/contact within 24h)
	// User 4: 2 points (example.com/ and example.com/blog)
	expected := map[string]int{
		"1": 2,
		"2": 3,
		"3": 1,
		"4": 2,
	}

	assert.Equal(t, expected, userPoints, "Points calculation incorrect")
}
