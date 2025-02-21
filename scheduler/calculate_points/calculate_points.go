package calculate_points

import (
	"context"
	"fmt"
	"log/slog"

	db "github.com/devs-group/driplet"
	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Run() error {
	slog.Info("calculating points...")

	sql, err := db.Connect(db.DefaultConfig())
	if err != nil {
		return errors.Wrap(err, "unable to connect to database")
	}
	defer sql.Close()

	points, err := CalculatePoints(context.Background(), sql)
	if err != nil {
		return errors.Wrap(err, "error calculating points")
	}

	// now write the points to the postgres DB by increasing the users credits
	tx, err := sql.BeginTxx(context.Background(), nil)
	if err != nil {
		return errors.Wrap(err, "unable to start transaction")
	}

	defer func() {
		if err != nil {
			// If there was an error, roll back the transaction
			_ = tx.Rollback()
		}
	}()

	// For each user with points, update their credits
	for userID, pointsEarned := range points {
		if pointsEarned <= 0 {
			continue // Skip users with no points
		}

		// Update the credits for this user
		_, err = tx.Exec(`
			UPDATE users
			SET credits = credits + $1
			WHERE id = $2
		`, pointsEarned, userID)

		if err != nil {
			return errors.Wrapf(err, "failed to update credits for user %d", userID)
		}

		slog.Info("updated credits",
			"user_id", userID,
			"points_earned", pointsEarned)
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	slog.Info("successfully updated credits for all users",
		"total_users_updated", len(points))

	return nil
}

// CalculatePoints calculates points for each user based on unique website visits in the last 24 hours
// Returns a map with user_id as key and points as value
func CalculatePoints(ctx context.Context, db *sqlx.DB) (map[int]int, error) {
	// SQL query to get unique website visits in the last 24 hours
	query := `
		WITH unique_visits AS (
			SELECT DISTINCT user_id, host, path
			FROM website_visits
			WHERE timestamp >= NOW() - INTERVAL '24 hours'
		)
		SELECT user_id, COUNT(*) as points
		FROM unique_visits
		GROUP BY user_id
	`

	type UserPoints struct {
		UserID int `db:"user_id"`
		Points int `db:"points"`
	}

	var results []UserPoints
	if err := db.SelectContext(ctx, &results, query); err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	userPoints := make(map[int]int)
	for _, row := range results {
		userPoints[row.UserID] = row.Points
	}

	return userPoints, nil
}
