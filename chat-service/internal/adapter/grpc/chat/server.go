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
	var c *client
	var roomIdx uint
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		switch payload := req.GetType().(type) {
		case *pb.ChatReq_Join:
			if err = s.join(payload.Join, c, stream); err != nil {
				return err
			}
			roomIdx = uint(payload.Join.RoomIdx)
			defer func() {
				if err = s.app.Disconnect(uint(payload.Join.RoomIdx), c); err != nil {
					s.logger.Errorf("disconnect client failed: %s", err)
				}
			}()
		case *pb.ChatReq_Message:
			if err = s.handleMessage(payload.Message); err != nil {
				return err
			}
		default:
			return errors.New("not available parameter")
		}
	}
}

func (s *Server) join(payload *pb.JoinReq, c *client, stream pb.Chat_ConnectServer) error {
	if c != nil {
		return errors.New("already joined")
	}
	// TODO: check user validation from user service
	c = &client{userIdx: uint(payload.UserIdx), send: stream.Send}
	s.app.Connect(uint(payload.RoomIdx), c)
	return nil
}

func (s *Server) handleMessage(payload *pb.MessageReq, roomIdx uint, c *client) error {
	return s.app.SendMessge(&dto.MessageDto{
		RoomIdx:  roomIdx,
		UserIdx:  c.userIdx,
		Name:     c.Name,
		ImageUrl: c.ImageUrl,
		Message:  payload.Message,
	})
}
