package port

import "github.com/widcraft/chat-service/internal/domain/entity"

type ChatApp interface {
	CreateRoom(name string) error
	ConnectRoom(roomId, userId uint) error
	JoinRoom(roomId uint) error
	SendMessge(roomId uint) error
}

type ChatRepository interface {
	CreateRoom(*entity.Room) error
	JoinRoom(roomId, userId uint) error
	SaveMessage() error
	GetMessage() (*entity.Message, error)
}
