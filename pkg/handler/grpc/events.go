package grpc

import (
	"context"
	"github.com/AssylzhanZharzhanov/dochq-test-task/gen/proto"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/service/grpc"
)

type EventsHandler struct {
	proto.UnimplementedEventsServiceServer

	service *grpc.Service
}

func NewEventsHandler(service *grpc.Service) *EventsHandler {
	return &EventsHandler{service: service}
}

func (h *EventsHandler) GetEvents(ctx context.Context, req *proto.GetEventsRequest) (*proto.GetEventsResponse, error) {

	return nil, nil
}