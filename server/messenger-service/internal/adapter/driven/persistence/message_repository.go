package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"messenger-service/internal/adapter/driven/persistence/db"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/domain/entity"
)

type MessageRepository struct {
	collection *mongo.Collection
}

func NewMessageRepository(mongoDb *db.MongoDb) *MessageRepository {
	collection := mongoDb.Database("chat").Collection("message")

	return &MessageRepository{collection}
}

func (mr MessageRepository) SaveMessage(message *entity.Message) error {
	if _, err := mr.collection.InsertOne(context.TODO(), message); err != nil {
		return err
	}

	return nil
}

func (mr MessageRepository) GetMessages(query *dto.MessagesQuery) ([]*entity.Message, error) {
	return nil, nil
}
