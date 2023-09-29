package repository

import (
	"github.com/widcraft/chat-service/internal/domain/entity"
	"github.com/widcraft/chat-service/internal/port"
	"github.com/widcraft/chat-service/pkg/db"
	"github.com/widcraft/chat-service/pkg/logger"
)

type ChatRepository struct {
	logger logger.Logger
	db     *db.Redis
}

func NewChatRepository(logger logger.Logger, redis *db.Redis) *ChatRepository {
	return &ChatRepository{
		logger: logger,
		db:     redis,
	}
}

func (repo *ChatRepository) ConnectRoom(roomIdx, client port.ChatClient) error {
	return nil
}

func (repo *ChatRepository) SaveMessage(message *entity.Message) error {
	return nil
}

func (repo *ChatRepository) GetMessages(roomIdx uint) ([]entity.Message, error) {
	return nil, nil
}
