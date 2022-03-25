package grpc

import (
	"context"
	"errors"
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/repository/grpc"
	"log"
)

const (
	createEvent = "create"
	updateEvent = "update"
	deleteEvent = "delete"
)

type AnswersService struct {
	answersRepo grpc.Answers
	eventsRepo grpc.Events
}

func (s *AnswersService) CreateAnswer(ctx context.Context, answer models.Answer) (models.Answer, error) {
	data, _ := s.answersRepo.GetAnswer(ctx, answer.Key)
	if data.Key != "" && data.Val != "" {
		return models.Answer{}, errors.New("this answer already exists")
	}

	output, err := s.answersRepo.CreateAnswer(ctx, answer)
	if err != nil {
		return output, err
	}

	event := models.Event{
		Event: createEvent,
		Data:  output,
	}
	err = s.eventsRepo.CreateEvent(ctx, event)
	if err != nil {
		return models.Answer{}, err
	}

	return output, nil
}

func (s *AnswersService) GetAnswer(ctx context.Context, key string) (models.Answer, error) {
	return s.answersRepo.GetAnswer(ctx, key)
}

func (s *AnswersService) UpdateAnswer(ctx context.Context, key string, value string) (models.Answer, error) {
	answer, err := s.answersRepo.UpdateAnswer(ctx, key, value)
	if (answer.Key == "" && answer.Val == "") || err != nil {
		log.Printf("key: %s", answer.Key)
		return models.Answer{}, errors.New("not found")
	}

	event := models.Event{
		Event: updateEvent,
		Data:  answer,
	}

	err = s.eventsRepo.CreateEvent(ctx, event)
	if err != nil {
		return models.Answer{}, err
	}

	return answer, nil}

func (s *AnswersService) DeleteAnswer(ctx context.Context, key string) error {
	answer, err := s.answersRepo.GetAnswer(ctx, key)
	if err != nil {
		return errors.New("not found")
	}

	event := models.Event{
		Event: deleteEvent,
		Data:  answer,
	}

	err = s.eventsRepo.CreateEvent(ctx, event)
	if err != nil {
		return err
	}

	return s.answersRepo.DeleteAnswer(ctx, key)}

func NewAnswersService(repository *grpc.Repository) *AnswersService {
	return &AnswersService{
		answersRepo: repository.Answers,
		eventsRepo: repository.Events,
	}
}
