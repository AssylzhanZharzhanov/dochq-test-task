package rest

import (
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

func (r *EventsRepository) CreateEvent(event models.Event) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (event, data)
		VALUES ($1, $2)
	`, eventsTable)

	_, err := r.db.Exec(query, event.Event, event.Data)
	return err
}

func (r *EventsRepository) GetEvents(key string) ([]models.EventOutput, error) {
	events := make([]models.EventOutput, 0)

	query := fmt.Sprintf(`
		SELECT 
			event,
			data
		FROM %s
		WHERE data->>'key' = $1
		ORDER BY id ASC
	`, eventsTable)

	err := r.db.Select(&events, query, key)
	return events, err
}

func (r *EventsRepository) GetLastEvent(key string) (models.Event, error) {
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

	err := r.db.Get(&event, query, key)
	return event, err
}

func NewEventsRepository(db *sqlx.DB) *EventsRepository {
	return &EventsRepository{db: db}
}

