package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"messenger-service/internal/adapter/driven/persistence/db"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/domain/entity"
	"messenger-service/internal/port/driven"
	"time"
)

type MessageRepository struct {
	logger     driven.Logger
	collection *mongo.Collection
}

func NewMessageRepository(logger driven.Logger, mongoDb *db.MongoDb) *MessageRepository {
	collection := mongoDb.Database("chat").Collection("message")

	return &MessageRepository{logger, collection}
}

func (mr *MessageRepository) SaveMessage(message *entity.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	if _, err := mr.collection.InsertOne(ctx, message); err != nil {
		return err
	}

	return nil
}

func (mr *MessageRepository) GetMessages(query *dto.MessagesQuery) ([]*entity.Message, error) {
	var messages []*entity.Message

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	filter := bson.M{}
	cursor, err := mr.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
