package singlemutex

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/port/driver"
	"sync"
)

type RoomManager struct {
	roomById map[string]*LiveRoom
	mu       sync.RWMutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		roomById: make(map[string]*LiveRoom),
	}
}

func (rm *RoomManager) Join(client driver.MessengerClient) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	roomId := client.RoomId()
	room, ok := rm.roomById[roomId]

	if !ok {
		room = NewLiveRoom(roomId)
		rm.roomById[roomId] = room
	}

	room.Join(client)
}

func (rm *RoomManager) Leave(client driver.MessengerClient) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	clientId := client.RoomId()
	room := rm.roomById[clientId]
	room.Leave(client)

	if room.IsEmpty() {
		delete(rm.roomById, clientId)
	}
}

func (rm *RoomManager) Broadcast(message dto.Message) error {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	room := rm.roomById[message.RoomId]

	if room == nil {
		return nil
	}

	return room.Broadcast(message)
}
