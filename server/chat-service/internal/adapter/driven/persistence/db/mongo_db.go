package db

import (
	"chat-service/internal/adapter/driven/config"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	*mongo.Client
}

func NewMongoDb(configStore *config.Config) (*MongoDb, error) {
	configs := configStore.GetMongoDbConfig()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", configs.User, configs.Password, configs.Host, configs.Port)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &MongoDb{client}, nil
}

func (md *MongoDb) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return md.Disconnect(ctx)
}
