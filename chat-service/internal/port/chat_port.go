package port

import (
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/domain/entity"
)

type ChatClient interface {
	GetUserIdx() uint32
	SendMessage(message *dto.MessageDto) error
}

type ChatApp interface {
	Connect(roomIdx uint32, client ChatClient)
	Disconnect(roomIdx uint32, client ChatClient) error
	SendMessge(messageDto *dto.MessageDto) error
	GetMessages(roomIdx uint32) ([]dto.MessageDto, error)
}

type ChatRepository interface {
	ConnectRoom(roomIdx, client ChatClient) error
	SaveMessage(message *entity.Message) error
	GetMessages(roomIdx uint32) ([]entity.Message, error)
}
