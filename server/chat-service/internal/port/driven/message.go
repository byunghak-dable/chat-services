package driven

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/entity"
)

type MessageRepository interface {
	Save(message *entity.Message) error
	GetMulti(query dto.MessagesQuery) ([]*entity.Message, error)
}
