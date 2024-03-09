package persistence

import (
	"messenger-service/internal/application/dto"
	"messenger-service/internal/domain/entity"
)

type MessageRepository interface {
	SaveMessage(message *entity.Message) error
	GetMessages(query *dto.MessagesQuery) ([]*entity.Message, error)
}
