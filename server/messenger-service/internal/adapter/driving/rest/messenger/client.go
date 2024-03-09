package messenger

import (
	"github.com/gorilla/websocket"
	"messenger-service/internal/application/dto"
)

type client struct {
	websocketConn *websocket.Conn
	userIdx       uint
	roomIdx       uint
}

func (c *client) SendMessage(message *dto.Message) error {
	return c.websocketConn.WriteJSON(message)
}

func (c *client) GetRoomIdx() uint {
	return c.roomIdx
}

func (c *client) GetUserIdx() uint {
	return c.userIdx
}
