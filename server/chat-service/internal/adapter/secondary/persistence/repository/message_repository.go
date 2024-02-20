package repository

import (
	"github.com/widcraft/chat-service/internal/adapter/secondary/persistence/db"
	"github.com/widcraft/chat-service/internal/domain/entity"
	"github.com/widcraft/chat-service/internal/port/secondary"
)

type MessageRepository struct {
	logger  secondary.Logger
	mongoDb *db.MongoDb
}

func NewMessageRepository(logger secondary.Logger, mongoDb *db.MongoDb) *MessageRepository {
	return &MessageRepository{
		logger:  logger,
		mongoDb: mongoDb,
	}
}

func (repo *MessageRepository) SaveMessage(message *entity.Message) error {
	return nil
}

func (repo *MessageRepository) GetMessages(roomIdx uint) ([]entity.Message, error) {
	return nil, nil
}
