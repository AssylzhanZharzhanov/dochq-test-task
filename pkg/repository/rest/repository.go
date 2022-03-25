package rest

import (
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/jmoiron/sqlx"
)

type Answers interface {
	CreateAnswer(answer models.Answer) (models.Answer, error)
	GetAnswer(key string) (models.Answer, error)
	UpdateAnswer(key string, value string) (models.Answer, error)
	DeleteAnswer(key string) error
}

type Events interface {
	CreateEvent(event models.Event) error
	GetEvents(key string) ([]models.EventOutput, error)
	GetLastEvent(key string) (models.Event, error)
}

type Repository struct {
	Answers
	Events
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Answers: NewAnswerRepository(db),
		Events:  NewEventsRepository(db),
	}
}
