package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/adapter/repository/redis"
	"github.com/widcraft/chat-service/internal/domain/entity"
	"github.com/widcraft/chat-service/internal/port"
)

type ChatRepository struct {
	logger *log.Logger
	db     *redis.Redis
}

func NewChatRepository(logger *log.Logger, redis *redis.Redis) *ChatRepository {
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
