package chat

import (
	"errors"
	"messenger-service/internal/adapter/driving/grpc/chat/pb"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
)

type Handler struct {
	pb.UnimplementedChatServer
	logger           driven.Logger
	messengerService driving.Messenger
}

func New(logger driven.Logger, app driving.Messenger) *Handler {
	return &Handler{
		logger:           logger,
		messengerService: app,
	}
}

func (h *Handler) Connect(stream pb.Chat_ConnectServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	joinReq, ok := req.GetType().(*pb.ChatReq_Join)
	if !ok {
		return errors.New("should join before other request")
	}

	return h.handleConnection(stream, joinReq.Join)
}

func (h *Handler) handleConnection(stream pb.Chat_ConnectServer, joinReq *pb.JoinReq) error {
	client := &client{
		stream:  stream,
		roomIdx: uint(joinReq.UserIdx),
		userIdx: uint(joinReq.RoomIdx),
	}

	if err := h.messengerService.Join(client); err != nil {
		return err
	}

	defer func() {
		if err := h.messengerService.Leave(client); err != nil {
			h.logger.Errorf("rest chat leave error: %v", err)
		}
	}()

	return h.handleMessage(client)
}

func (h *Handler) handleMessage(client *client) error {
	for {
		req, err := client.stream.Recv()
		if err != nil {
			return err
		}

		switch typedReq := req.GetType().(type) {
		case *pb.ChatReq_Message:
			h.sendMessage(client, typedReq.Message)
		default:
			return errors.New("wrong request type")
		}
	}
}

func (h *Handler) sendMessage(client *client, payload *pb.MessageReq) {
	message := dto.Message{
		RoomIdx:  client.roomIdx,
		UserIdx:  client.userIdx,
		Name:     client.name,
		ImageUrl: client.imageUrl,
		Message:  payload.GetMessage(),
	}

	if err := h.messengerService.SendMessage(&message); err != nil {
		h.logger.Errorf("send message failed: %h", err)
	}
}
