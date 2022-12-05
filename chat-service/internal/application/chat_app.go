package application

import (
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

func (app *ChatApp) Connect(roomIdx uint, client port.ChatClient) {
	app.chatPool.RegisterJob(func() {
		room, ok := app.rooms[roomIdx]
		if ok {
			app.rooms[roomIdx] = append(room, client)
			return
		}
		app.rooms[roomIdx] = []port.ChatClient{client}
	})
}

func (app *ChatApp) Disconnect(roomIdx uint, client port.ChatClient) error {
	return nil
}

func (app *ChatApp) SendMessge(roomIdx uint, message dto.MessageDto) error {
	return nil
}

func (app *ChatApp) GetMessages(roomIdx uint) ([]dto.MessageDto, error) {
	return nil, nil
}
