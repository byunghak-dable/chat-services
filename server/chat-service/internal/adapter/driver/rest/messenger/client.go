package messenger

import (
	"chat-service/internal/application/dto"
	"github.com/gorilla/websocket"
)

type client struct {
	conn   *websocket.Conn
	userId string
	roomId string
}

func (c *client) Send(message dto.Message) error {
	return c.conn.WriteJSON(message)
}

func (c *client) RoomId() string {
	return c.roomId
}

func (c *client) UserId() string {
	return c.userId
}
