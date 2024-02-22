package chat

import (
	"github.com/gorilla/websocket"
	"github.com/widcraft/chat-service/internal/application/dto"
)

type client struct {
	websocketConn *websocket.Conn
	name          string
	imageUrl      string
	userIdx       uint
	roomIdx       uint
}

func (c *client) GetRoomIdx() uint {
	return c.roomIdx
}

func (c *client) GetUserIdx() uint {
	return c.userIdx
}

func (c *client) SendMessage(message *dto.MessageDto) error {
	return c.websocketConn.WriteJSON(message)
}
