package chat

import (
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/adapter/grpc/chat/pb"
)

type Server struct {
	logger *log.Logger
	pb.UnimplementedChatServer
}

func New(logger *log.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) Receive(req *pb.ConnectReq, stream pb.Chat_ReceiveServer) error {
	return nil
}

func (s *Server) Send(stream pb.Chat_SendServer) error {
	return nil
}
