package chat

import (
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/adapter/grpc/chat/pb"
	"github.com/widcraft/chat-service/internal/port"
)

type Server struct {
	logger *log.Logger
	app    port.ChatApp
	pb.UnimplementedChatServer
}

func New(logger *log.Logger, app port.ChatApp) *Server {
	return &Server{
		logger: logger,
		app:    app,
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
