package repository

import (
	"github.com/widcraft/chat-service/internal/domain/entity"
	"github.com/widcraft/chat-service/pkg/db"
	"github.com/widcraft/chat-service/pkg/logger"
)

type MessageRepository struct {
	logger logger.Logger
	db     *db.Redis
}

func NewMessageRepository(logger logger.Logger, redis *db.Redis) *MessageRepository {
	return &MessageRepository{
		logger: logger,
		db:     redis,
	}
}

func (repo *MessageRepository) SaveMessage(message *entity.Message) error {
	return nil
}

func (repo *MessageRepository) GetMessages(roomIdx uint) ([]entity.Message, error) {
	return nil, nil
}
