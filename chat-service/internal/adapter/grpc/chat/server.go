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

func (s *Server) Connect(stream pb.Chat_ConnectServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		switch v := req.GetType().(type) {
		case *pb.ChatReq_Join:
			s.logger.Infoln("join", v)
		case *pb.ChatReq_Message:
			s.logger.Infoln("message", v)
		}
	}
	return nil
}
