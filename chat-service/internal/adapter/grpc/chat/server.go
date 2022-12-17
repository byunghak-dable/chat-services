package chat

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/adapter/grpc/chat/pb"
	"github.com/widcraft/chat-service/internal/domain/dto"
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
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	joinReq, ok := req.GetType().(*pb.ChatReq_Join)
	if !ok {
		return errors.New("should join before other request")
	}

	userIdx, roomIdx := uint(joinReq.Join.RoomIdx), uint(joinReq.Join.UserIdx)
	c := &client{userIdx: userIdx, send: stream.Send}
	s.app.Connect(roomIdx, c)
	defer func() {
		if err = s.app.Disconnect(roomIdx, c); err != nil {
			s.logger.Errorf("disconnect client failed: %s", err)
		}
	}()
	return s.handleConnection(stream, roomIdx, c)
}

func (s *Server) handleConnection(stream pb.Chat_ConnectServer, roomIdx uint, c *client) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		switch payload := req.GetType().(type) {
		case *pb.ChatReq_Message:
			err = s.app.SendMessge(&dto.MessageDto{
				RoomIdx:  roomIdx,
				UserIdx:  c.userIdx,
				Name:     c.name,
				ImageUrl: c.imageUrl,
				Message:  payload.Message.GetMessage(),
			})
			if err != nil {
				s.logger.Errorf("send message failed: %s", err)
			}
		default:
			return errors.New("wrong request type")
		}
	}
}
