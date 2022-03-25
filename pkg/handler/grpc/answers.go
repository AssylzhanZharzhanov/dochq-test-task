package grpc

import (
	"context"
	"github.com/AssylzhanZharzhanov/dochq-test-task/gen/proto"
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/AssylzhanZharzhanov/dochq-test-task/pkg/service/grpc"
)

type AnswerHandler struct {
	proto.UnimplementedAnswersServiceServer

	service *grpc.Service
}

func (h *AnswerHandler) CreateAnswer(ctx context.Context, req *proto.Answer) (*proto.Answer, error) {
	input := models.Answer{
		Key: req.Key,
		Val: req.Value,
	}

	output, err := h.service.CreateAnswer(ctx, input)
	if err != nil {
		return nil, err
	}

	return toAnswerEntity(output), nil
}

func (h *AnswerHandler) GetAnswer(ctx context.Context, req *proto.GetAnswerRequest) (*proto.Answer, error) {
	key := req.Key

	output, err := h.service.GetAnswer(ctx, key)
	if err != nil {
		return nil, err
	}

	return toAnswerEntity(output),nil
}

func (h *AnswerHandler) UpdateAnswer(ctx context.Context, req *proto.Answer) (*proto.Answer, error) {
	input := &models.Answer{
		Key: req.Key,
		Val: req.Value,
	}

	output, err := h.service.UpdateAnswer(ctx, input.Key, input.Val)
	if err != nil {
		return nil, err
	}

	return toAnswerEntity(output), nil
}

func (h *AnswerHandler) DeleteAnswer(ctx context.Context, req *proto.DeleteAnswerRequest) (*proto.DeleteAnswerResponse, error) {

	key := req.Key

	err := h.service.DeleteAnswer(ctx, key)
	if err != nil {
		return nil, err
	}

	return &proto.DeleteAnswerResponse{Message: "DELETED"}, nil
}

func NewAnswerHandler(service *grpc.Service) *AnswerHandler {
	return &AnswerHandler{service: service}
}

func toAnswerEntity(answer models.Answer) *proto.Answer {
	return & proto.Answer{
		Key: answer.Key,
		Value: answer.Val,
	}
}

