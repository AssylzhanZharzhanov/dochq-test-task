package grpc

import (
	"fmt"
	"github.com/AssylzhanZharzhanov/dochq-test-task/gen/proto"
	grpcHandler "github.com/AssylzhanZharzhanov/dochq-test-task/pkg/handler/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
}

func NewServer() *Server {
	return &Server{
		grpcServer: grpc.NewServer(),
	}
}

func (s *Server) Run(port string, handler *grpcHandler.Handler) error {

	addr := fmt.Sprintf(":%s", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	proto.RegisterAnswersServiceServer(s.grpcServer, handler.AnswerHandler)
	proto.RegisterEventsServiceServer(s.grpcServer, handler.EventsHandler)

	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("%s", err.Error())
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}