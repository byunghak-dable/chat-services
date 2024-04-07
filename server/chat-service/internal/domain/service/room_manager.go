package service

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/entity"
	"chat-service/internal/port/driver"
	"sync"
	"sync/atomic"
)

type RoomManager struct {
	rooms   map[string]*entity.LiveRoom
	mu      sync.RWMutex
	tickets int32
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[string]*entity.LiveRoom),
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

	println(room)
	return room.Broadcast(message)
}

func (rm *RoomManager) getOrCreateRoom(roomId string) *entity.LiveRoom {
	if room := rm.getRoom(roomId); room != nil {
		return room
	}

	return rm.createRoom(roomId)
}

func (rm *RoomManager) getRoom(roomId string) *entity.LiveRoom {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	return rm.rooms[roomId]
}

func (rm *RoomManager) createRoom(roomId string) *entity.LiveRoom {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if room, ok := rm.rooms[roomId]; ok {
		return room
	}

	rm.rooms[roomId] = entity.NewLiveRoom(roomId)

	return rm.rooms[roomId]
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
	for id, room := range rm.rooms {
		if room.IsEmpty() {
			delete(rm.rooms, id)
		}
	}
}

func (rm *RoomManager) withTicket(action func()) {
	atomic.AddInt32(&rm.tickets, 1)
	defer atomic.AddInt32(&rm.tickets, -1)

	action()
}

func (rm *RoomManager) isCleanable() bool {
	return atomic.LoadInt32(&rm.tickets) == 1
}
