package chat

import (
	"errors"
	"messenger-service/internal/adapter/driving/grpc/chat/pb"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
)

type Server struct {
	pb.UnimplementedChatServer
	logger           driven.LoggerPort
	messengerService driving.MessengerServicePort
}

func New(logger driven.LoggerPort, app driving.MessengerServicePort) *Server {
	return &Server{
		logger:           logger,
		messengerService: app,
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

	return s.handleConnection(stream, joinReq.Join)
}

func (s *Server) handleConnection(stream pb.Chat_ConnectServer, joinReq *pb.JoinReq) error {
	client := &client{
		stream:  stream,
		roomIdx: uint(joinReq.UserIdx),
		userIdx: uint(joinReq.RoomIdx),
	}

	if err := s.messengerService.Join(client); err != nil {
		return err
	}

	defer func() {
		if err := s.messengerService.Leave(client); err != nil {
			s.logger.Errorf("rest chat leave error: %v", err)
		}
	}()

	return s.handleMessage(client)
}

func (s *Server) handleMessage(client *client) error {
	for {
		req, err := client.stream.Recv()
		if err != nil {
			return err
		}

		switch typedReq := req.GetType().(type) {
		case *pb.ChatReq_Message:
			s.sendMessage(client, typedReq.Message)
		default:
			return errors.New("wrong request type")
		}
	}
}

func (s *Server) sendMessage(client *client, payload *pb.MessageReq) {
	message := dto.MessageDto{
		RoomIdx:  client.roomIdx,
		UserIdx:  client.userIdx,
		Name:     client.name,
		ImageUrl: client.imageUrl,
		Message:  payload.GetMessage(),
	}

	if err := s.messengerService.SendMessage(&message); err != nil {
		s.logger.Errorf("send message failed: %s", err)
	}
}
