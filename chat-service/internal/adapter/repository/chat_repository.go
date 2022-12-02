package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/adapter/repository/redis"
	"github.com/widcraft/chat-service/internal/domain/entity"
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

func (repo *ChatRepository) CreateRoom(room *entity.Room) error {
	return nil
}

func (repo *ChatRepository) JoinRoom(roomIdx, userIdx uint) error {
	return nil
}

func (repo *ChatRepository) SaveMessage(userIdx uint, message string) error {
	return nil
}

func (repo *ChatRepository) GetMessages(roomIdx uint) ([]entity.Message, error) {
	return nil, nil
}
