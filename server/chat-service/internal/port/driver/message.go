package driver

import "chat-service/internal/application/dto"

type Message interface {
	Save(message *dto.Message) error
	GetMulti(query dto.MessagesQuery) ([]dto.Message, error)
}
