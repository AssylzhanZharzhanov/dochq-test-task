package grpc

import (
	"context"
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/repository/grpc"
)

type EventsService struct {
	repos grpc.Events
}

func (s *EventsService) GetEvents(ctx context.Context, key string) ([]models.EventOutput, error) {
	return s.repos.GetEvents(ctx, key)
}

func NewEventsService(repos *grpc.Repository) *EventsService {
	return &EventsService{repos: repos.Events}
}
