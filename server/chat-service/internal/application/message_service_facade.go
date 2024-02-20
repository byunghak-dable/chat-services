package application

import (
	"github.com/widcraft/chat-service/internal/application/dto"
	"github.com/widcraft/chat-service/internal/port/primary"
	"github.com/widcraft/chat-service/internal/port/secondary"
)

type MessageService interface {
	SaveMessage(message *dto.MessageDto) error
	GetMessages(roomIdx uint) ([]dto.MessageDto, error)
}

type MessengerService interface {
	Participate(client primary.MessengerClient)
	Quit(client primary.MessengerClient)
	SendMessage(message *dto.MessageDto) error
}

type ChatService struct {
	logger           secondary.Logger
	messageService   MessageService
	messengerService MessengerService
}

func NewChatService(logger secondary.Logger, messageService MessageService, messengerService MessengerService) *ChatService {
	return &ChatService{
		logger:           logger,
		messageService:   messageService,
		messengerService: messengerService,
	}
}

func (c *ChatService) Join(client primary.MessengerClient) {
	c.messengerService.Participate(client)
}

func (c *ChatService) Leave(client primary.MessengerClient) {
	c.messengerService.Quit(client)
}

func (c *ChatService) SendMessage(message *dto.MessageDto) error {
	err := c.messengerService.SendMessage(message)
	if err != nil {
		return err
	}

	return c.messageService.SaveMessage(message)
}

func (c *ChatService) GetMessages(roomIdx uint) ([]dto.MessageDto, error) {
	return nil, nil
}
