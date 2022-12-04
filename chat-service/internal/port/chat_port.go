package port

import (
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/domain/entity"
)

type ChatClient interface {
	GetUseridx() uint
	SendMessage() error
}

type ChatApp interface {
	ConnectRoom(roomIdx, client ChatClient) error
	SendMessge(roomIdx uint, messageDto dto.MessageDto) error
	GetMessages(roomIdx uint) ([]dto.MessageDto, error)
	Disconnect(client ChatClient) error
}

type ChatRepository interface {
	ConnectRoom(roomIdx, client ChatClient) error
	SaveMessage(roomIdx, userIdx uint, message string) error
	GetMessages(roomIdx uint) ([]entity.Message, error)
}
