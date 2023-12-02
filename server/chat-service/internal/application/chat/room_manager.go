package chat

import (
	"bytes"
	"errors"
	"fmt"
	"sync"

	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/port"
	"github.com/widcraft/chat-service/pkg/logger"
)

type roomManager struct {
	logger logger.Logger
	rooms  map[uint][]port.ChatClient
	mutex  *sync.RWMutex
}

func NewRoomManager(logger logger.Logger) *roomManager {
	return &roomManager{
		logger: logger,
		rooms:  make(map[uint][]port.ChatClient),
		mutex:  new(sync.RWMutex),
	}
}

func (manager *roomManager) add(client port.ChatClient) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	roomIdx := client.GetRoomIdx()
	room, ok := manager.rooms[roomIdx]

	if ok {
		manager.rooms[roomIdx] = append(room, client)
		manager.logger.Infof("room id: %d, count: %d", roomIdx, len(room))
		return
	}

	manager.rooms[roomIdx] = []port.ChatClient{client}
	manager.logger.Infof("room id: %d added", roomIdx, len(room))
}

func (manager *roomManager) quit(client port.ChatClient) {
	// TODO: try to find better way to handle race condition than mutex
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	roomIdx := client.GetRoomIdx()
	room, ok := manager.rooms[roomIdx]

	if !ok {
		manager.logger.Error("no existing chat room roomIdx")
		return
	}

	for i, participant := range room {
		if client == participant {
			manager.rooms[roomIdx] = append(room[:i], room[i+1:]...)
		}
	}

	manager.logger.Error("no client in chat room")
}

func (manager *roomManager) sendMessage(message *dto.MessageDto) error {
	manager.mutex.RLock()
	defer manager.mutex.RUnlock()

	room, ok := manager.rooms[message.RoomIdx]

	if !ok {
		return errors.New("no existing chat room roomIdx")
	}

	var errBuffer bytes.Buffer

	for _, client := range room {
		err := client.SendMessage(message)
		if err != nil {
			errBuffer.WriteString(fmt.Sprintf("client %d send failed: %s\n", client.GetUserIdx(), err.Error()))
		}
	}

	if errBuffer.Len() == 0 {
		return errors.New(errBuffer.String())
	}

	return nil
}
