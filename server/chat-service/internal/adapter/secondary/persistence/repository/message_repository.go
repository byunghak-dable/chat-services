package repository

import (
	"github.com/widcraft/chat-service/internal/domain/entity"
	"github.com/widcraft/chat-service/internal/port/secondary"
)

type MessageDao interface {
	GetMessages(roomIdx string) ([]*entity.Message, error)

	SaveMessage(msg *entity.Message) error
}

type MessageRepository struct {
	logger     secondary.Logger
	messageDao MessageDao
}

func NewMessageRepository(logger secondary.Logger, messageDao MessageDao) *MessageRepository {
	return &MessageRepository{
		logger:     logger,
		messageDao: messageDao,
	}
}

func (repo *MessageRepository) SaveMessage(message *entity.Message) error {
	return nil
}

func (repo *MessageRepository) GetMessages(roomIdx uint) ([]entity.Message, error) {
	return nil, nil
}
