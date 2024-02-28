package driven

import (
	"encoding/json"
	"fmt"
	"messenger-service/internal/application/dto"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	AcksNone = 0
	AcksOne  = 1
	AcksAll  = "all"
)

type KafkaNode struct {
	Host string
	Port string
}

type KafkaProducerConfig struct {
	ClientId string
	Topic    string
	Nodes    []KafkaNode
}

type KafkaProducer struct {
	*kafka.Producer
	topic string
}

func NewKafkaProducer(config *KafkaProducerConfig) (*KafkaProducer, error) {
	var servers []string
	for _, node := range config.Nodes {
		servers = append(servers, fmt.Sprintf("%s:%s", node.Host, node.Port))
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": strings.Join(servers, ","),
		"client.id":         config.ClientId,
		"acks":              AcksAll,
	})
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{producer, config.Topic}, nil
}

func (producer *KafkaProducer) ProduceMessage(message *dto.MessageDto) error {
	kafkaMessage, err := producer.makeMessage(message)
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	err = producer.Produce(kafkaMessage, nil)
	if err != nil {
		return fmt.Errorf("failed to produce message: %v", err)
	}

	switch event := (<-producer.Events()).(type) {
	case *kafka.Message:
		if event.TopicPartition.Error != nil {
			return fmt.Errorf("delivery failed: %v", event.TopicPartition.Error)
		}
	default:
		return fmt.Errorf("unknown event: %v", event)
	}

	return nil
}

func (producer *KafkaProducer) makeMessage(message *dto.MessageDto) (*kafka.Message, error) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	return &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &producer.topic, Partition: kafka.PartitionAny},
		Value:          []byte(jsonMessage),
	}, nil
}

func (producer *KafkaProducer) OnExit() error {
	producer.Close()
	return nil
}
