package application

import (
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/port"
)

type user interface {
	GetIdx() uint
	GetName() string
	GetImageUrl() string
	Send() error
}

type workerPool interface {
	RegisterJob(func())
}

type ChatApp struct {
	logger   *log.Logger
	pool     workerPool
	chatRoom map[int][]user
	repo     port.ChatRepository
}

func NewChatApp(logger *log.Logger, pool workerPool, repo port.ChatRepository) *ChatApp {
	return &ChatApp{
		logger:   logger,
		pool:     pool,
		chatRoom: make(map[int][]user),
		repo:     repo,
	}
}

func (app *ChatApp) Connect(roomIdx uint, client port.ChatClient) error {
	return nil
}

func (app *ChatApp) SendMessge(roomIdx uint, message dto.MessageDto) error {
	return nil
}

func (app *ChatApp) GetMessages(roomIdx uint) ([]dto.MessageDto, error) {
	return nil, nil
}

func (app *ChatApp) Disconnect(roomIdx uint, client port.ChatClient) error {
	return nil
}
