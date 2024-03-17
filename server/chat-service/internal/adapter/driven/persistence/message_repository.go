package persistence

import (
	"chat-service/internal/adapter/driven/persistence/db"
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/entity"
	"chat-service/internal/port/driven"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (mr *MessageRepository) Save(message *entity.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	result, err := mr.collection.InsertOne(ctx, message)
	if err != nil {
		return err
	}

	message.SetId(result.InsertedID.(primitive.ObjectID).Hex())
	return nil
}

func (mr *MessageRepository) GetSeveral(query dto.MessagesQuery) ([]*entity.Message, error) {
	var messages []*entity.Message

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	filter := options.FindOptions{}
	cursor, err := mr.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
