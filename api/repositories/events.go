package repositories

import "github.com/jmoiron/sqlx"

type EventsRepository struct {
	DB *sqlx.DB
}

func NewEventsRepository(db *sqlx.DB) (*EventsRepository, error) {
	return &EventsRepository{DB: db}, nil
}
