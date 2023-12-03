package storage

import (
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/domain/entity"
	"github.com/widcraft/chat-service/internal/port"
	"github.com/widcraft/chat-service/pkg/logger"
)

type ChatStorageApp struct {
	logger logger.Logger
	repo   port.ChatRepository
}

func New(logger logger.Logger, repo port.ChatRepository) *ChatStorageApp {
	return &ChatStorageApp{
		logger: logger,
		repo:   repo,
	}
}

func (app *ChatStorageApp) SaveMessage(message *dto.MessageDto) error {
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

func (app *ChatStorageApp) GetMessages(roomIdx uint) ([]dto.MessageDto, error) {
	return nil, nil
}
