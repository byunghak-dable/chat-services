package message

import (
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/domain/entity"
	"github.com/widcraft/chat-service/internal/infra"
	"github.com/widcraft/chat-service/internal/port"
)

type MessageStorageService struct {
	logger infra.Logger
	repo   port.MessageRepository
}

func NewMessageStorageService(logger infra.Logger, repo port.MessageRepository) *MessageStorageService {
	return &MessageStorageService{
		logger: logger,
		repo:   repo,
	}
}

func (app *MessageStorageService) SaveMessage(message *dto.MessageDto) error {
	err := app.repo.SaveMessage(&entity.Message{
		RoomIdx:  message.RoomIdx,
		UserIdx:  message.UserIdx,
		ImageUrl: message.ImageUrl,
		Name:     message.Name,
		Message:  message.Message,
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *MessageStorageService) GetMessages(roomIdx uint) ([]dto.MessageDto, error) {
	return nil, nil
}
