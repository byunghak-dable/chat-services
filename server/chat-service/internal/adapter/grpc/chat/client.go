package chat

import (
	"github.com/widcraft/chat-service/internal/adapter/grpc/chat/pb"
	"github.com/widcraft/chat-service/internal/domain/dto"
)

type client struct {
	send     func(*pb.MessageRes) error
	name     string
	imageUrl string
	userIdx  uint
}

func (c *client) GetUserIdx() uint {
	return c.userIdx
}

func (c *client) SendMessage(message *dto.MessageDto) error {
	return c.send(&pb.MessageRes{
		RoomIdx:  uint32(message.RoomIdx),
		UserIdx:  uint32(message.UserIdx),
		Message:  message.Message,
		Name:     message.Name,
		ImageUrl: message.ImageUrl,
	})
}
