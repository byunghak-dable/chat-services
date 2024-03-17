package driven

import "chat-service/internal/application/dto"

type MessageSubscriber interface {
	OnReceive(message *dto.Message)
}

type MessageBroker interface {
	Subscribe(subscriber MessageSubscriber)
	Publish(message *dto.Message) error
}
