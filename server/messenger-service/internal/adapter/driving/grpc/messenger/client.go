package messenger

import (
	"messenger-service/internal/adapter/driving/grpc/messenger/pb"
	"messenger-service/internal/application/dto"
)

type client struct {
	stream  pb.Chat_ConnectServer
	roomIdx uint
	userIdx uint
}

func (c *client) SendMessage(message *dto.Message) error {
	messageRes := pb.MessageRes{
		RoomIdx: uint32(message.RoomIdx),
		UserIdx: uint32(message.UserIdx),
		Message: message.Message,
	}

	return c.stream.Send(&messageRes)
}

func (c *client) GetRoomIdx() uint {
	return c.roomIdx
}

func (c *client) GetUserIdx() uint {
	return c.userIdx
}
