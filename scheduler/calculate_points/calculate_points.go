package calculate_points

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/devs-group/driplet/pkg/db"
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

	points, err := CalculatePoints(context.Background(), sql.SQLX)
	if err != nil {
		return errors.Wrap(err, "error calculating points")
	}

	tx, err := sql.SQLX.BeginTxx(context.Background(), nil)
	if err != nil {
		return errors.Wrap(err, "unable to start transaction")
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// for each user with points, update their credits
	for userID, pointsEarned := range points {
		if pointsEarned <= 0 {
			continue // skip users with no points
		}

		_, err = tx.Exec(`
			UPDATE users
			SET credits = credits + $1
			WHERE id = $2
		`, pointsEarned, userID)
		if err != nil {
			return errors.Wrapf(err, "failed to update credits for user %s", userID)
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
func CalculatePoints(ctx context.Context, db *sqlx.DB) (map[string]int, error) {
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
		UserID string `db:"user_id"`
		Points int    `db:"points"`
	}

	var results []UserPoints
	if err := db.SelectContext(ctx, &results, query); err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	userPoints := make(map[string]int)
	for _, row := range results {
		userPoints[row.UserID] = row.Points
	}

	return userPoints, nil
}
