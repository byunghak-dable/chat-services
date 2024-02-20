package dao

import (
	"context"

	"github.com/widcraft/chat-service/internal/adapter/secondary/persistence/client"
	"github.com/widcraft/chat-service/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageDao struct {
	collection *mongo.Collection
}

func NewMessageDao(mongoDb *client.MongoDb) *MessageDao {
	return &MessageDao{
		collection: mongoDb.Database("chat").Collection("message"),
	}
}

func (dao *MessageDao) GetMessages(roomIdx string) ([]*entity.Message, error) {
	return nil, nil
}

func (dao *MessageDao) SaveMessage(message *entity.Message) error {
	_, err := dao.collection.InsertOne(context.TODO(), message)

	return err
}
