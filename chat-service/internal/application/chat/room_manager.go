package chat

import (
	"bytes"
	"errors"
	"fmt"
	"sync"

	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/port"
)

type roomManager struct {
	rooms map[uint][]port.ChatClient
	mutex *sync.RWMutex
}

func NewRoomManager() *roomManager {
	return &roomManager{
		rooms: make(map[uint][]port.ChatClient),
		mutex: new(sync.RWMutex),
	}
}

func (manager *roomManager) add(roomIdx uint, client port.ChatClient) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	room, ok := manager.rooms[roomIdx]
	if ok {
		manager.rooms[roomIdx] = append(room, client)
		return
	}
	manager.rooms[roomIdx] = []port.ChatClient{client}
}

func (manager *roomManager) quit(roomIdx uint, client port.ChatClient) error {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	room, ok := manager.rooms[roomIdx]
	if !ok {
		return errors.New("no existing chat room roomIdx")
	}
	for i, roomClient := range room {
		if client == roomClient {
			manager.rooms[roomIdx] = append(room[:i], room[i+1:]...)
			return nil
		}
	}
	return errors.New("no client in chat room")
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
		if err := client.SendMessage(message); err != nil {
			errBuffer.WriteString(fmt.Sprintf("client %d send failed: %s\n", client.GetUserIdx(), err.Error()))
		}
	}

	if errBuffer.Len() != 0 {
		return errors.New(errBuffer.String())
	}
	return nil
}
