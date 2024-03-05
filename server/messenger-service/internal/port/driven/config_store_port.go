package driven

type ConfigStore interface {
	GetKafkaServers() string
	GetKafkaGroupId() string
	GetKafkaClientId() string
	GetKafkaChatTopic() string
	GetRestPort() string
	GetGrpcPort() string
}
