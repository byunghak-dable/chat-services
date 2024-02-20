package message

import (
	"github.com/widcraft/chat-service/internal/application/dto"
	"github.com/widcraft/chat-service/internal/domain/entity"
	"github.com/widcraft/chat-service/internal/port/secondary"
)

type messageService struct {
	logger secondary.Logger
	repo   secondary.MessageRepository
}

func NewMessageService(logger secondary.Logger, repo secondary.MessageRepository) *messageService {
	return &messageService{
		logger: logger,
		repo:   repo,
	}
}

func (app *messageService) SaveMessage(message *dto.MessageDto) error {
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

func (app *messageService) GetMessages(roomIdx uint) ([]dto.MessageDto, error) {
	return nil, nil
}
