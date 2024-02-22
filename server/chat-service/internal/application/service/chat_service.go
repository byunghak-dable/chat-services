package service

import (
	"github.com/widcraft/chat-service/internal/application/abstraction"
	"github.com/widcraft/chat-service/internal/application/dto"
	"github.com/widcraft/chat-service/internal/port/driven"
	"github.com/widcraft/chat-service/internal/port/driving"
)

type ChatService struct {
	logger           driven.Logger
	messageService   abstraction.MessageService
	messengerService abstraction.MessengerService
}

func NewChatService(
	logger driven.Logger,
	messageService abstraction.MessageService,
	messengerService abstraction.MessengerService,
) *ChatService {
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
