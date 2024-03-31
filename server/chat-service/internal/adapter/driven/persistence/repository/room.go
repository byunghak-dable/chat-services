package repository

import (
	"chat-service/internal/adapter/driven/persistence/db"
	"chat-service/internal/domain/entity"
	"chat-service/internal/port/driven"

	"go.mongodb.org/mongo-driver/mongo"
)

type Room struct {
	logger     driven.Logger
	collection *mongo.Collection
}

func NewRoom(logger driven.Logger, mongoDb *db.MongoDb) *Room {
	collection := mongoDb.Database("chat").Collection("room")

	return &Room{logger, collection}
}

func (r *Room) Save(room *entity.Message) {
}

func (r *Room) GetMulti() {
}
