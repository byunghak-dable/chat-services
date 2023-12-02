package chat

import (
	"github.com/widcraft/chat-service/internal/adapter/grpc/chat/pb"
	"github.com/widcraft/chat-service/internal/domain/dto"
)

type client struct {
	stream   pb.Chat_ConnectServer
	name     string
	imageUrl string
	roomIdx  uint
	userIdx  uint
}

func (c *client) SendMessage(message *dto.MessageDto) error {
	return c.stream.Send(&pb.MessageRes{
		RoomIdx:  uint32(message.RoomIdx),
		UserIdx:  uint32(message.UserIdx),
		Message:  message.Message,
		Name:     message.Name,
		ImageUrl: message.ImageUrl,
	})
}

func (c *client) GetRoomIdx() uint {
	return c.roomIdx
}

func (c *client) GetUserIdx() uint {
	return c.userIdx
}
