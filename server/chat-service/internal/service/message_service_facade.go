package service

import (
	"github.com/widcraft/chat-service/internal/adapter/secondary/repository"
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/infra"
	"github.com/widcraft/chat-service/internal/port"
	"github.com/widcraft/chat-service/internal/service/message"
)

type MessageServiceFacade struct {
	logger           infra.Logger
	storageService   *message.StorageService
	messengerService *message.MessengerService
}

func NewMessageServiceFacade(logger infra.Logger, messageRepo *repository.MessageRepository) *MessageServiceFacade {
	return &MessageServiceFacade{
		logger:           logger,
		storageService:   message.NewStorageService(logger, messageRepo),
		messengerService: message.NewMessengerService(logger),
	}
}

func (facade *MessageServiceFacade) Join(client port.MessengerClient) {
	facade.messengerService.Participate(client)
}

func (facade *MessageServiceFacade) Leave(client port.MessengerClient) {
	facade.messengerService.Quit(client)
}

func (facade *MessageServiceFacade) SendMessge(message *dto.MessageDto) error {
	err := facade.messengerService.SendMessage(message)
	if err != nil {
		return err
	}

	return facade.storageService.SaveMessage(message)
}

func (facade *MessageServiceFacade) GetMessages(roomIdx uint) ([]dto.MessageDto, error) {
	return nil, nil
}
