package chat

import (
	"github.com/widcraft/chat-service/internal/application/chat/messenger"
	"github.com/widcraft/chat-service/internal/application/chat/storage"
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/port"
	"github.com/widcraft/chat-service/pkg/logger"
)

type ChatFacade struct {
	logger       logger.Logger
	storageApp   *storage.ChatStorageApp
	messengerApp *messenger.ChatMessengerApp
}

func New(logger logger.Logger, storageApp *storage.ChatStorageApp, messengerApp *messenger.ChatMessengerApp) *ChatFacade {
	return &ChatFacade{logger, storageApp, messengerApp}
}

func (facade *ChatFacade) Connect(client port.ChatClient) {
	facade.messengerApp.Participate(client)
}

func (facade *ChatFacade) Disconnect(client port.ChatClient) {
	facade.messengerApp.Quit(client)
}

func (facade *ChatFacade) SendMessge(message *dto.MessageDto) error {
	err := facade.messengerApp.SendMessage(message)
	if err != nil {
		return err
	}

	return facade.storageApp.SaveMessage(message)
}

func (facade *ChatFacade) GetMessages(roomIdx uint) ([]dto.MessageDto, error) {
	return nil, nil
}
