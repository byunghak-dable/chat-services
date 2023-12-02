package port

import (
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/domain/entity"
)

type ChatClient interface {
	GetRoomIdx() uint
	GetUserIdx() uint
	SendMessage(message *dto.MessageDto) error
}

type ChatApp interface {
	Connect(client ChatClient)
	Disconnect(client ChatClient)
	SendMessge(messageDto *dto.MessageDto) error
	GetMessages(roomIdx uint) ([]dto.MessageDto, error)
}

type ChatRepository interface {
	ConnectRoom(roomIdx, client ChatClient) error
	SaveMessage(message *entity.Message) error
	GetMessages(roomIdx uint) ([]entity.Message, error)
}
