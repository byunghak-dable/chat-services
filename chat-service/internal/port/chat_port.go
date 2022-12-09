package port

import (
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/domain/entity"
)

type ChatClient interface {
	GetUserIdx() uint
	SendMessage(message interface{}) error
}

type ChatApp interface {
	Connect(roomIdx uint, client ChatClient) error
	Disconnect(roomIdx uint, client ChatClient) error
	SendMessge(messageDto *dto.MessageDto) error
	GetMessages(roomIdx uint) ([]dto.MessageDto, error)
}

type ChatRepository interface {
	ConnectRoom(roomIdx, client ChatClient) error
	SaveMessage(message *entity.Message) error
	GetMessages(roomIdx uint) ([]entity.Message, error)
}
