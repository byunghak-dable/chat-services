package chat

import (
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/domain/entity"
	"github.com/widcraft/chat-service/internal/port"
	"github.com/widcraft/chat-service/pkg/logger"
)

type ChatApp struct {
	logger     logger.Logger
	roomManger *roomManager
	repo       port.ChatRepository
}

func New(logger logger.Logger, repo port.ChatRepository) *ChatApp {
	return &ChatApp{
		logger:     logger,
		roomManger: NewRoomManager(logger),
		repo:       repo,
	}
}

func (app *ChatApp) Connect(client port.ChatClient) {
	app.roomManger.add(client)
}

func (app *ChatApp) Disconnect(client port.ChatClient) {
	app.roomManger.quit(client)
}

func (app *ChatApp) SendMessge(message *dto.MessageDto) error {
	// TODO: save message
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

	return app.roomManger.sendMessage(message)
}

func (app *ChatApp) GetMessages(roomIdx uint) ([]dto.MessageDto, error) {
	return nil, nil
}
