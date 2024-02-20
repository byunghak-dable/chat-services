package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
}

type MongoDb struct {
	client *mongo.Client
}

func NewMongoDb(config MongoDbConfig) (*MongoDb, error) {
	applyUri := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.User, config.User, config.Host, config.Port)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(applyUri))
	if err != nil {
		return nil, err
	}

	return &MongoDb{client: client}, nil
}
