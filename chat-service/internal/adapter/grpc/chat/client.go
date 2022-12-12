package chat

import (
	"github.com/widcraft/chat-service/internal/adapter/grpc/chat/pb"
	"github.com/widcraft/chat-service/internal/domain/dto"
)

type client struct {
	userIdx uint
	send    func(*pb.MessageRes) error
}

func (c *client) GetUserIdx() uint {
	return c.userIdx
}

func (c *client) SendMessage(message *dto.MessageDto) error {
	return c.send(&pb.MessageRes{
		RoomIdx:  message.RoomIdx,
		UserIdx:  message.UserIdx,
		Name:     message.Name,
		ImageUrl: message.ImageUrl,
	})
}
