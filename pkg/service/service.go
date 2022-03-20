package service

import (
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/repository"
)

type Answers interface {
	CreateAnswer(answer models.Answer) (models.Answer, error)
	GetAnswer(key string) (models.Answer, error)
	UpdateAnswer(key string, value string) (models.Answer, error)
	DeleteAnswer(key string) error
}

type Events interface {
	GetEvents(key string) ([]models.EventOutput, error)
}

type Service struct {
	Answers
	Events
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Answers: NewAnswerService(repository),
		Events: NewEventsService(repository),
	}
}
