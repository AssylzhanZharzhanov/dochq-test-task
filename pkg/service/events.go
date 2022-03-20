package service

import (
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/repository"
)

type EventsService struct {
	repos repository.Events
}

func (s *EventsService) GetEvents(key string) ([]models.EventOutput, error) {
	return s.repos.GetEvents(key)
}

func NewEventsService(repos *repository.Repository) *EventsService {
	return &EventsService{repos: repos.Events}
}
