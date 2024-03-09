package config

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type Store struct {
}

func NewStore() (*Store, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Store{}, nil
}

func (cs *Store) GetMongoDbConfig() MongoDb {
	return MongoDb{
		Host:     os.Getenv("MONGODB_HOST"),
		Port:     os.Getenv("MONGODB_PORT"),
		User:     os.Getenv("MONGODB_USER"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		Db:       os.Getenv("MONGODB_DB"),
	}
}

func (cs *Store) GetKafkaProducerConfig() KafkaProducer {
	return KafkaProducer{
		kafkaBase: cs.getKafkaConfig(),
		ClientId:  os.Getenv("KAFKA_CLIENT_ID"),
	}
}

func (cs *Store) GetKafkaConsumerConfig() KafkaConsumer {
	return KafkaConsumer{
		kafkaBase: cs.getKafkaConfig(),
		GroupId:   fmt.Sprintf("%s:%s", os.Getenv("KAFKA_GROUP_ID"), uuid.New()),
	}
}

func (cs *Store) getKafkaConfig() kafkaBase {
	servers := []string{
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_1_HOST"), os.Getenv("KAFKA_1_PORT")),
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_2_HOST"), os.Getenv("KAFKA_2_PORT")),
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_3_HOST"), os.Getenv("KAFKA_3_PORT")),
	}

	return kafkaBase{
		Servers: strings.Join(servers, ","),
		Topic:   os.Getenv("KAFKA_CHAT_TOPIC"),
	}
}

func (cs *Store) GetRestPort() string {
	return os.Getenv("REST_PORT")
}

func (cs *Store) GetGrpcPort() string {
	return os.Getenv("GRPC_PORT")
}
