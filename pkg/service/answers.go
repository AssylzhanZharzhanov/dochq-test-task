package service

import (
	"errors"
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/repository"
	"log"
)

const (
	createEvent = "create"
	updateEvent = "update"
	deleteEvent = "delete"
)

type AnswerService struct {
	repos repository.Answers
	eventsRepos repository.Events
}

func (s *AnswerService) CreateAnswer(answer models.Answer) (models.Answer, error) {
	data, _ := s.repos.GetAnswer(answer.Key)

	if data.Key != "" && data.Val != "" {
		return models.Answer{}, errors.New("this answer already exists")
	}

	output, err := s.repos.CreateAnswer(answer)
	if err != nil {
		return models.Answer{}, err
	}

	event := models.Event{
		Event: createEvent,
		Data: output,
	}
	err = s.eventsRepos.CreateEvent(event)
	if err != nil {
		return models.Answer{}, err
	}

	return output, nil
}

func (s *AnswerService) GetAnswer(key string) (models.Answer, error) {
	return s.repos.GetAnswer(key)
}

func (s *AnswerService) UpdateAnswer(key string, value string) (models.Answer, error) {
	answer, err := s.repos.UpdateAnswer(key, value)
	if (answer.Key == "" && answer.Val == "") || err != nil {
		log.Printf("key: %s", answer.Key)
		return models.Answer{}, errors.New("not found")
	}

	event := models.Event{
		Event: updateEvent,
		Data: answer,
	}

	err = s.eventsRepos.CreateEvent(event)
	if err != nil {
		return models.Answer{}, err
	}

	return answer, nil
}

func (s *AnswerService) DeleteAnswer(key string) error {
	answer, err := s.repos.GetAnswer(key)
	if err != nil {
		return errors.New("not found")
	}

	event := models.Event{
		Event: deleteEvent,
		Data: answer,
	}

	err = s.eventsRepos.CreateEvent(event)
	if err != nil {
		return err
	}

	return s.repos.DeleteAnswer(key)
}

func NewAnswerService(repository *repository.Repository) *AnswerService {
	return &AnswerService{
		repos: repository.Answers,
		eventsRepos: repository.Events,
	}
}