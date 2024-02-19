package repository

import (
	"github.com/widcraft/chat-service/internal/domain/entity"
	"github.com/widcraft/chat-service/internal/infra"
	"github.com/widcraft/chat-service/internal/infra/db"
)

type MessageRepository struct {
	logger  infra.Logger
	mongoDb *db.MongoDb
}

func NewMessageRepository(logger infra.Logger, mongoDb *db.MongoDb) *MessageRepository {
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
