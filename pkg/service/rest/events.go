package rest

import (
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/repository/rest"
)

type EventsService struct {
	repos rest.Events
}

func (s *EventsService) GetEvents(key string) ([]models.EventOutput, error) {
	return s.repos.GetEvents(key)
}

func NewEventsService(repos *rest.Repository) *EventsService {
	return &EventsService{repos: repos.Events}
}
