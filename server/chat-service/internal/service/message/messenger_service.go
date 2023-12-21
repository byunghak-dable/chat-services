package message

import (
	"fmt"
	"sync"

	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/infra"
	"github.com/widcraft/chat-service/internal/port"
)

type MessengerService struct {
	logger infra.Logger
	rooms  map[uint][]port.MessengerClient
	mutex  *sync.RWMutex
}

func NewMessengerService(logger infra.Logger) *MessengerService {
	return &MessengerService{
		logger: logger,
		rooms:  make(map[uint][]port.MessengerClient),
		mutex:  new(sync.RWMutex),
	}
}

func (service *MessengerService) Participate(client port.MessengerClient) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	roomIdx := client.GetRoomIdx()
	room, ok := service.rooms[roomIdx]

	if ok {
		service.rooms[roomIdx] = append(room, client)
		service.logger.Infof("room id: %d, count: %d", roomIdx, len(room))
		return
	}

	service.rooms[roomIdx] = []port.MessengerClient{client}
	service.logger.Infof("room id: %d added", roomIdx, len(room))
}

func (app *MessengerService) Quit(client port.MessengerClient) {
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

func (service *MessengerService) SendMessage(message *dto.MessageDto) error {
	service.mutex.RLock()
	defer service.mutex.RUnlock()

	room, ok := service.rooms[message.RoomIdx]

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
