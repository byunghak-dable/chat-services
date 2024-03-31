package service

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/entity"
	"chat-service/internal/port/driver"
	"sync"
)

type RoomManager struct {
	roomById map[string]*entity.LiveRoom
	lock     sync.RWMutex
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		roomById: make(map[string]*entity.LiveRoom),
	}
}

func (rm *RoomManager) Join(client driver.MessengerClient) {
	rm.getRoom(client.RoomId()).Join(client)
}

func (rm *RoomManager) Leave(client driver.MessengerClient) {
	roomId := client.RoomId()
	room := rm.getRoom(roomId)

	room.Leave(client)
	rm.withLock(func() {
		if room.IsEmpty() {
			delete(rm.roomById, roomId)
		}
	})
}

func (rm *RoomManager) Broadcast(message dto.Message) error {
	return rm.getRoom(message.RoomId).Broadcast(message)
}

func (rm *RoomManager) getRoom(roomId string) *entity.LiveRoom {
	var room *entity.LiveRoom

	rm.withRLock(func() {
		room = rm.roomById[roomId]
	})

	if room != nil {
		return room
	}

	rm.withLock(func() {
		room = rm.roomById[roomId]

		if room == nil {
			room = entity.NewLiveRoom(roomId)
			rm.roomById[roomId] = room
		}
	})

	return room
}

func (rm *RoomManager) withLock(action func()) {
	rm.lock.Lock()
	defer rm.lock.Unlock()

	action()
}

func (rm *RoomManager) withRLock(action func()) {
	rm.lock.RLock()
	defer rm.lock.RUnlock()

	action()
}
