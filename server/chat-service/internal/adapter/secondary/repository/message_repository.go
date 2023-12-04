package repository

import (
	"github.com/widcraft/chat-service/internal/adapter/secondary/repository/db"
	"github.com/widcraft/chat-service/internal/domain/entity"
	"github.com/widcraft/chat-service/internal/infra"
)

type MessageRepository struct {
	logger infra.Logger
	db     *db.Redis
}

func NewMessageRepository(logger infra.Logger, redis *db.Redis) *MessageRepository {
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
