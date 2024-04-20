package repository

import (
	"chat-service/internal/adapter/driven/persistence/db"
	"chat-service/internal/adapter/driven/persistence/orm"
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/entity"
	"chat-service/internal/port/driven"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Message struct {
	logger     driven.Logger
	collection *mongo.Collection
}

func NewMessage(logger driven.Logger, mongoDb *db.MongoDb) *Message {
	collection := mongoDb.Database("chat").Collection("message")

	return &Message{logger, collection}
}

func (m *Message) Create(message *entity.Message) error {
	ctx := context.TODO()

	result, err := m.collection.InsertOne(ctx, orm.FromMessage(*message))
	if err != nil {
		return err
	}

	message.SetId(result.InsertedID.(primitive.ObjectID).Hex())

	return nil
}

func (m *Message) GetMulti(query dto.MessagesQuery) ([]entity.Message, error) {
	ctx := context.TODO()
	filter := []bson.E{{Key: "room_id", Value: query.RoomId}}
	findOpts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}).SetLimit(query.Limit)

	if cursorId, err := primitive.ObjectIDFromHex(query.Cursor); err == nil {
		filter = append(filter, bson.E{Key: "_id", Value: bson.D{{Key: "$gte", Value: cursorId}}})
	}

	cursor, err := m.collection.Find(ctx, filter, findOpts)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = cursor.Close(ctx)
	}()

	var messages []entity.Message

	for cursor.Next(ctx) {
		var message orm.Message

		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}

		messages = append(messages, message.ToDomain())
	}

	return messages, nil
}
