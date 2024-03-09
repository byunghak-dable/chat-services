package driven

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type ConfigStore struct {
}

func NewConfigStore() (*ConfigStore, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &ConfigStore{}, nil
}

func (cs *ConfigStore) GetKafkaServers() string {
	servers := []string{
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_1_HOST"), os.Getenv("KAFKA_1_PORT")),
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_2_HOST"), os.Getenv("KAFKA_2_PORT")),
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_3_HOST"), os.Getenv("KAFKA_3_PORT")),
	}

	return strings.Join(servers, ",")
}

func (cs *ConfigStore) GetKafkaGroupId() string {
	return fmt.Sprintf("%s:%s", os.Getenv("KAFKA_GROUP_ID"), uuid.New())
}

func (cs *ConfigStore) GetKafkaClientId() string {
	return os.Getenv("KAFKA_CLIENT_ID")
}

func (cs *ConfigStore) GetKafkaChatTopic() string {
	return os.Getenv("KAFKA_CHAT_TOPIC")
}

func (cs *ConfigStore) GetRestPort() string {
	return os.Getenv("REST_PORT")
}

func (cs *ConfigStore) GetGrpcPort() string {
	return os.Getenv("GRPC_PORT")
}
