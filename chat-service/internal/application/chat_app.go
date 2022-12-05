package application

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/port"
)

type workerPool interface {
	RegisterJob(func())
}

type ChatApp struct {
	logger   *log.Logger
	chatPool workerPool
	rooms    map[uint][]port.ChatClient
	repo     port.ChatRepository
}

func NewChatApp(logger *log.Logger, pool workerPool, repo port.ChatRepository) *ChatApp {
	return &ChatApp{
		logger:   logger,
		chatPool: pool,
		rooms:    make(map[uint][]port.ChatClient),
		repo:     repo,
	}
}

func (app *ChatApp) Connect(roomIdx uint, client port.ChatClient) error {
	// TODO: check if user is valid from user service
	app.chatPool.RegisterJob(func() {
		room, ok := app.rooms[roomIdx]
		if ok {
			app.rooms[roomIdx] = append(room, client)
			return
		}
		app.rooms[roomIdx] = []port.ChatClient{client}
	})
	return nil
}

func (app *ChatApp) Disconnect(roomIdx uint, client port.ChatClient) error {
	errChan := make(chan error)
	defer close(errChan)

	app.chatPool.RegisterJob(func() {
		room, ok := app.rooms[roomIdx]
		if !ok {
			errChan <- errors.New("no existing chat room roomIdx")
			return
		}
		for i, roomClient := range room {
			if client == roomClient {
				app.rooms[roomIdx] = append(room[:i], room[i+1:]...)
				return
			}
		}
		errChan <- errors.New("no client in chat room")
	})
	return <-errChan
}

func (app *ChatApp) SendMessge(roomIdx uint, message dto.MessageDto) error {
	return nil
}

func (app *ChatApp) GetMessages(roomIdx uint) ([]dto.MessageDto, error) {
	return nil, nil
}
