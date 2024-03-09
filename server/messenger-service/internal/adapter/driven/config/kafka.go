package config

type KafkaProducer struct {
	kafkaBase
	ClientId string
}

type KafkaConsumer struct {
	kafkaBase
	GroupId string
}

type kafkaBase struct {
	Servers string
	Topic   string
}
