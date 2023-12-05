package port

import (
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/domain/entity"
)

type MessengerClient interface {
	GetRoomIdx() uint
	GetUserIdx() uint
	SendMessage(message *dto.MessageDto) error
}

type MessageService interface {
	Join(client MessengerClient)
	Leave(client MessengerClient)
	SendMessge(messageDto *dto.MessageDto) error
	GetMessages(roomIdx uint) ([]dto.MessageDto, error)
}

type MessageRepository interface {
	SaveMessage(message *entity.Message) error
	GetMessages(roomIdx uint) ([]entity.Message, error)
}
