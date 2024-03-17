package driven

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/entity"
)

type MessageRepository interface {
	Save(message *entity.Message) error
	GetSeveral(query *dto.MessagesQuery) ([]*entity.Message, error)
}

type MessageSubscriber interface {
	OnReceive(message *dto.Message)
}

type MessageBroker interface {
	Subscribe(subscriber MessageSubscriber)
	Publish(message *dto.Message) error
}
