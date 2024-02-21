package service

import (
	"github.com/widcraft/chat-service/internal/application/dto"
	"github.com/widcraft/chat-service/internal/port/driven"
	"github.com/widcraft/chat-service/internal/port/driving"
)

type MessageService interface {
	SaveMessage(message *dto.MessageDto) error
	GetMessages(roomIdx uint) ([]dto.MessageDto, error)
}

type MessengerService interface {
	Participate(client driving.MessengerClient)
	Quit(client driving.MessengerClient)
	SendMessage(message *dto.MessageDto) error
}

type ChatService struct {
	logger           driven.Logger
	messageService   MessageService
	messengerService MessengerService
}

func NewChatService(logger driven.Logger, messageService MessageService, messengerService MessengerService) *ChatService {
	return &ChatService{
		logger:           logger,
		messageService:   messageService,
		messengerService: messengerService,
	}
}

func (c *ChatService) Join(client driving.MessengerClient) {
	c.messengerService.Participate(client)
}

func (c *ChatService) Leave(client driving.MessengerClient) {
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
