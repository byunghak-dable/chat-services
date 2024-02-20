package secondary

import "github.com/widcraft/chat-service/internal/domain/entity"

type MessageRepository interface {
	SaveMessage(message *entity.Message) error
	GetMessages(roomIdx uint) ([]entity.Message, error)
}
