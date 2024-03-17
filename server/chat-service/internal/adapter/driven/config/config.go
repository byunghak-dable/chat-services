package config

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type MongoDb struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
}

type MessageBroker struct {
	Servers  string
	Topic    string
	GroupId  string
	ClientId string
}

type Config struct {
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{}, nil
}

func (cs *Config) GetMongoDbConfig() MongoDb {
	return MongoDb{
		Host:     os.Getenv("MONGODB_HOST"),
		Port:     os.Getenv("MONGODB_PORT"),
		User:     os.Getenv("MONGODB_USER"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		Db:       os.Getenv("MONGODB_DB"),
	}
}

func (cs *Config) GetMessageBrokerConfig() MessageBroker {
	servers := []string{
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_1_HOST"), os.Getenv("KAFKA_1_PORT")),
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_2_HOST"), os.Getenv("KAFKA_2_PORT")),
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_3_HOST"), os.Getenv("KAFKA_3_PORT")),
	}

	return MessageBroker{
		Servers:  strings.Join(servers, ","),
		Topic:    os.Getenv("KAFKA_CHAT_TOPIC"),
		ClientId: os.Getenv("KAFKA_CLIENT_ID"),
		GroupId:  fmt.Sprintf("%s:%s", os.Getenv("KAFKA_GROUP_ID"), uuid.New()),
	}
}

func (cs *Config) GetRestPort() string {
	return os.Getenv("REST_PORT")
}

func (cs *Config) GetGrpcPort() string {
	return os.Getenv("GRPC_PORT")
}
