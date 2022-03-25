package grpc

import (
	"context"
	"fmt"
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/jmoiron/sqlx"
)

const (
	eventsTable = "events"
)

type EventsRepository struct {
	db *sqlx.DB
}

func (r *EventsRepository) CreateEvent(ctx context.Context, event models.Event) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (event, data)
		VALUES ($1, $2)
	`, eventsTable)

	_, err := r.db.ExecContext(ctx, query, event.Event, event.Data)
	return err
}

func (r *EventsRepository) GetEvents(ctx context.Context, key string) ([]models.EventOutput, error) {
	events := make([]models.EventOutput, 0)

	query := fmt.Sprintf(`
		SELECT 
			event,
			data
		FROM %s
		WHERE data->>'key' = $1
		ORDER BY id ASC
	`, eventsTable)

	err := r.db.SelectContext(ctx, &events, query, key)
	return events, err
}

func (r *EventsRepository) GetLastEvent(ctx context.Context, key string) (models.Event, error) {
	event := models.Event{}

	query := fmt.Sprintf(`
		SELECT 
			event, 
			data
		FROM %s
		WHERE data->>'key' = $1
		ORDER BY id DESC
		LIMIT 1
	`, eventsTable)

	err := r.db.GetContext(ctx, &event, query, key)
	return event, err
}

func NewEventsRepository(db *sqlx.DB) *EventsRepository {
	return &EventsRepository{db: db}
}
