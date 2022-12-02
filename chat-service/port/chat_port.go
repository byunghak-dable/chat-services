package port

import (
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/domain/entity"
)

type ChatApp interface {
	CreateRoom(name string) (uint, error)
	ConnectRoom(roomIdx, userIdx uint) error
	JoinRoom(roomIdx uint) error
	SendMessge(roomIdx uint, message string) error
	GetMessages(roomIdx uint) ([]dto.MessageDto, error)
}

type ChatRepository interface {
	CreateRoom(room *entity.Room) error
	JoinRoom(roomIdx, userIdx uint) error
	SaveMessage(userIdx uint, message string) error
	GetMessages(roomIdx uint) ([]entity.Message, error)
}
