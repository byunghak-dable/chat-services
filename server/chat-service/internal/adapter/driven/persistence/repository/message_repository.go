package repository

import (
	"github.com/widcraft/chat-service/internal/adapter/driven/persistence/client"
	"github.com/widcraft/chat-service/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageDao interface {
	GetMessages(roomIdx string) ([]*entity.Message, error)

	SaveMessage(msg *entity.Message) error
}

type MessageRepository struct {
	messageCollection *mongo.Collection
}

func NewMessageRepository(mongoDb *client.MongoDb) *MessageRepository {
	return &MessageRepository{
		messageCollection: mongoDb.Database("chat").Collection("message"),
	}
}

func (repo *MessageRepository) SaveMessage(message *entity.Message) error {
	return nil
}

func (repo *MessageRepository) GetMessages(roomIdx uint) ([]entity.Message, error) {
	return nil, nil
}
