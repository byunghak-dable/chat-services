package driver

import "chat-service/internal/application/dto"

type GetMultiMessageUseCase interface {
	Handle(message dto.MessagesQuery) ([]dto.Message, error)
}
