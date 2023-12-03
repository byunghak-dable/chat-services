package messenger

import (
	"fmt"
	"sync"

	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/port"
	"github.com/widcraft/chat-service/pkg/logger"
)

type ChatMessengerApp struct {
	logger logger.Logger
	rooms  map[uint][]port.ChatClient
	mutex  *sync.RWMutex
}

func New(logger logger.Logger) *ChatMessengerApp {
	return &ChatMessengerApp{
		logger: logger,
		rooms:  make(map[uint][]port.ChatClient),
		mutex:  new(sync.RWMutex),
	}
}

func (app *ChatMessengerApp) Participate(client port.ChatClient) {
	app.mutex.Lock()
	defer app.mutex.Unlock()

	roomIdx := client.GetRoomIdx()
	room, ok := app.rooms[roomIdx]

	if ok {
		app.rooms[roomIdx] = append(room, client)
		app.logger.Infof("room id: %d, count: %d", roomIdx, len(room))
		return
	}

	app.rooms[roomIdx] = []port.ChatClient{client}
	app.logger.Infof("room id: %d added", roomIdx, len(room))
}

func (app *ChatMessengerApp) Quit(client port.ChatClient) {
	// TODO: try to find better way to handle race condition than mutex
	app.mutex.Lock()
	defer app.mutex.Unlock()

	roomIdx := client.GetRoomIdx()
	room, ok := app.rooms[roomIdx]

	if !ok {
		app.logger.Error("no existing chat room roomIdx")
		return
	}

	for i, participant := range room {
		if client == participant {
			app.rooms[roomIdx] = append(room[:i], room[i+1:]...)
		}
	}

	app.logger.Error("no client in chat room")
}

func (app *ChatMessengerApp) SendMessage(message *dto.MessageDto) error {
	app.mutex.RLock()
	defer app.mutex.RUnlock()

	room, ok := app.rooms[message.RoomIdx]

	if !ok {
		return fmt.Errorf("no existing chat room roomIdx")
	}

	sendErrors := []error{}

	for _, client := range room {
		err := client.SendMessage(message)
		if err != nil {
			sendErrors = append(sendErrors, err)
		}
	}

	if len(sendErrors) > 0 {
		return fmt.Errorf("send message errors, %v", sendErrors)
	}

	return nil
}
