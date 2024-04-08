package driver

import "chat-service/internal/application/dto"

type ReadMultiMessageUseCase interface {
	Handle(message dto.MessagesQuery) ([]dto.Message, error)
}
