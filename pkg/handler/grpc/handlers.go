package grpc

import (
	"github.com/AssylzhanZharzhanov/dochq-test-task/gen/proto"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/service/grpc"
)

type Handler struct {
	AnswerHandler proto.AnswersServiceServer
	EventsHandler proto.EventsServiceServer
}

func NewHandler(service *grpc.Service) *Handler {
	return &Handler{
		AnswerHandler: NewAnswerHandler(service),
		EventsHandler: NewEventsHandler(service),
	}
}