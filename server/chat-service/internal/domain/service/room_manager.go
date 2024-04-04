package service

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/entity"
	"chat-service/internal/port/driver"
	"sync"
	"sync/atomic"
)

type RoomManager struct {
	roomById map[string]*entity.LiveRoom
	mu       sync.RWMutex
	ticket   int32
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		roomById: make(map[string]*entity.LiveRoom),
	}
}

func (rm *RoomManager) Join(client driver.MessengerClient) {
	rm.withTicket(func() {
		rm.getOrCreateRoom(client.RoomId()).Join(client)
	})
}

func (rm *RoomManager) Leave(client driver.MessengerClient) {
	roomId := client.RoomId()

	rm.withTicket(func() {
		room := rm.getRoom(roomId)
		room.Leave(client)

		if room.IsEmpty() {
			rm.cleanRooms()
		}
	})
}

func (rm *RoomManager) Broadcast(message dto.Message) error {
	room := rm.getRoom(message.RoomId)

	if room == nil {
		return nil
	}

	return room.Broadcast(message)
}

func (rm *RoomManager) getOrCreateRoom(roomId string) *entity.LiveRoom {
	room := rm.getRoom(roomId)

	if room != nil {
		return room
	}

	return rm.createRoom(roomId)
}

func (rm *RoomManager) getRoom(roomId string) *entity.LiveRoom {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	return rm.roomById[roomId]
}

func (rm *RoomManager) createRoom(roomId string) *entity.LiveRoom {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	room, ok := rm.roomById[roomId]

	if !ok {
		room = entity.NewLiveRoom(roomId)
		rm.roomById[roomId] = room
	}

	return room
}

func (rm *RoomManager) cleanRooms() {
	if !rm.isCleanable() {
		return
	}

	rm.mu.Lock()
	defer rm.mu.Unlock()

	if !rm.isCleanable() {
		return
	}

	// TODO: need optimizing to only remove empty room
	for id, room := range rm.roomById {
		if room.IsEmpty() {
			delete(rm.roomById, id)
		}
	}
}

func (rm *RoomManager) withTicket(action func()) {
	atomic.AddInt32(&rm.ticket, 1)
	defer atomic.AddInt32(&rm.ticket, -1)

	action()
}

func (rm *RoomManager) isCleanable() bool {
	return atomic.LoadInt32(&rm.ticket) == 1
}
