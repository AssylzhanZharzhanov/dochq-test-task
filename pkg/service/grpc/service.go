package grpc

import (
	"context"
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/repository/grpc"
)

type Answers interface {
	CreateAnswer(ctx context.Context, answer models.Answer) (models.Answer, error)
	GetAnswer(ctx context.Context, key string) (models.Answer, error)
	UpdateAnswer(ctx context.Context, key string, value string) (models.Answer, error)
	DeleteAnswer(ctx context.Context, key string) error
}

type Events interface {
	GetEvents(ctx context.Context, key string) ([]models.EventOutput, error)
}

type Service struct {
	Answers
	Events
}

func NewService(repository *grpc.Repository) *Service {
	return &Service{
		Answers: NewAnswersService(repository),
		Events: NewEventsService(repository),
	}
}

