package application

import (
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/port"
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
	poolChan chan error
	chatRoom map[int][]user
	repo     port.ChatRepository
}

func NewChatApp(logger *log.Logger, pool workerPool, repo port.ChatRepository) *ChatApp {
	return &ChatApp{
		logger:   logger,
		pool:     pool,
		poolChan: make(chan error),
		chatRoom: make(map[int][]user),
		repo:     repo,
	}
}

func (app *ChatApp) CreateRoom(name string) (uint, error) {
	return 0, nil
}

func (app *ChatApp) ConnectRoom(roomIdx, userIdx uint) error {
	return nil
}

func (app *ChatApp) JoinRoom(roomIdx uint) error {
	return nil
}

func (app *ChatApp) SendMessge(roomIdx uint, message string) error {
	return nil
}

func (app *ChatApp) GetMessages(roomIdx uint) ([]dto.MessageDto, error) {
	return nil, nil
}
