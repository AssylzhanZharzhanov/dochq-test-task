package grpc

import (
	"context"
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/jmoiron/sqlx"
)

type Answers interface {
	CreateAnswer(ctx context.Context, answer models.Answer) (models.Answer, error)
	GetAnswer(ctx context.Context, key string) (models.Answer, error)
	UpdateAnswer(ctx context.Context, key string, value string) (models.Answer, error)
	DeleteAnswer(ctx context.Context, key string) error
}

type Events interface {
	CreateEvent(ctx context.Context, event models.Event) error
	GetEvents(ctx context.Context, key string) ([]models.EventOutput, error)
	GetLastEvent(ctx context.Context, key string) (models.Event, error)
}

type Repository struct {
	Answers
	Events
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Answers: NewAnswerRepository(db),
		Events: NewEventsRepository(db),
	}
}
