package service

import (
	"github.com/widcraft/chat-service/internal/adapter/secondary/repository"
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/infra"
	"github.com/widcraft/chat-service/internal/port"
	"github.com/widcraft/chat-service/internal/service/message"
)

type MessageFacade struct {
	logger           infra.Logger
	storageService   *message.MessageStorageService
	messengerService *message.MessengerService
}

func NewMessageFacade(logger infra.Logger, messageRepo *repository.MessageRepository) *MessageFacade {
	return &MessageFacade{
		logger:           logger,
		storageService:   message.NewMessageStorageService(logger, messageRepo),
		messengerService: message.NewMessengerService(logger),
	}
}

func (facade *MessageFacade) Connect(client port.MessengerClient) {
	facade.messengerService.Participate(client)
}

func (facade *MessageFacade) Disconnect(client port.MessengerClient) {
	facade.messengerService.Quit(client)
}

func (facade *MessageFacade) SendMessge(message *dto.MessageDto) error {
	err := facade.messengerService.SendMessage(message)
	if err != nil {
		return err
	}

	return facade.storageService.SaveMessage(message)
}

func (facade *MessageFacade) GetMessages(roomIdx uint) ([]dto.MessageDto, error) {
	return nil, nil
}
