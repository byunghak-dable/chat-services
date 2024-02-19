package service

import (
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/infra"
	"github.com/widcraft/chat-service/internal/port"
)

type MessageService interface {
	SaveMessage(message *dto.MessageDto) error
	GetMessages(roomIdx uint) ([]dto.MessageDto, error)
}

type MessengerService interface {
	Participate(client port.MessengerClient)
	Quit(client port.MessengerClient)
	SendMessage(message *dto.MessageDto) error
}

type ChatService struct {
	logger           infra.Logger
	messageService   MessageService
	messengerService MessengerService
}

func NewChatService(logger infra.Logger, messageService MessageService, messengerService MessengerService) *ChatService {
	return &ChatService{
		logger:           logger,
		messageService:   messageService,
		messengerService: messengerService,
	}
}

func (c *ChatService) Join(client port.MessengerClient) {
	c.messengerService.Participate(client)
}

func (c *ChatService) Leave(client port.MessengerClient) {
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
