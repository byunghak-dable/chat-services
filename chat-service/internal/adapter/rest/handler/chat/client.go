package chat

import "github.com/widcraft/chat-service/internal/domain/dto"

type client struct {
	userIdx uint
	send    func(interface{}) error
}

func (c *client) GetUserIdx() uint {
	return c.userIdx
}

func (c *client) SendMessage(message *dto.MessageDto) error {
	return c.send(message)
}
