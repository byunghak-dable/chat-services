package message

import (
	"fmt"
	"sync"

	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/port"
	"github.com/widcraft/chat-service/pkg/logger"
)

type MessengerService struct {
	logger logger.Logger
	rooms  map[uint][]port.MessengerClient
	mutex  *sync.RWMutex
}

func NewMessengerService(logger logger.Logger) *MessengerService {
	return &MessengerService{
		logger: logger,
		rooms:  make(map[uint][]port.MessengerClient),
		mutex:  new(sync.RWMutex),
	}
}

func (app *MessengerService) Participate(client port.MessengerClient) {
	app.mutex.Lock()
	defer app.mutex.Unlock()

	roomIdx := client.GetRoomIdx()
	room, ok := app.rooms[roomIdx]

	if ok {
		app.rooms[roomIdx] = append(room, client)
		app.logger.Infof("room id: %d, count: %d", roomIdx, len(room))
		return
	}

	app.rooms[roomIdx] = []port.MessengerClient{client}
	app.logger.Infof("room id: %d added", roomIdx, len(room))
}

func (app *MessengerService) Quit(client port.MessengerClient) {
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

func (app *MessengerService) SendMessage(message *dto.MessageDto) error {
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
