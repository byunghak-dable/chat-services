package service

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/entity"
	"chat-service/internal/port/driver"
	"sync"
)

type RoomManager struct {
	roomById map[string]*entity.LiveRoom
	lockById *sync.Map
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		roomById: make(map[string]*entity.LiveRoom),
		lockById: new(sync.Map),
	}
}

func (rm *RoomManager) Join(client driver.MessengerClient) {
	roomId := client.RoomId()

	rm.withLock(roomId, func() {
		if rm.roomById[roomId] == nil {
			rm.roomById[roomId] = entity.NewLiveRoom(roomId)
		}
		rm.roomById[roomId].Join(client)
	})
}

func (rm *RoomManager) Leave(client driver.MessengerClient) {
	roomId := client.RoomId()

	rm.withLock(roomId, func() {
		room := rm.roomById[roomId]
		room.Leave(client)

		if room.IsEmpty() {
			rm.clean(roomId)
		}
	})
}

func (rm *RoomManager) Broadcast(message dto.Message) error {
	var err error
	roomId := message.RoomId

	rm.withRLock(roomId, func() {
		err = rm.roomById[roomId].Broadcast(message)
	})

	return err
}

func (rm *RoomManager) clean(roomId string) {
	delete(rm.roomById, roomId)
	rm.lockById.Delete(roomId)
}

func (rm *RoomManager) withLock(roomId string, action func()) {
	lock := rm.getLock(roomId)
	lock.Lock()
	defer lock.Unlock()

	action()
}

func (rm *RoomManager) withRLock(roomId string, action func()) {
	lock := rm.getLock(roomId)
	lock.RLock()
	defer lock.RUnlock()

	action()
}

func (rm *RoomManager) getLock(roomId string) *sync.RWMutex {
	lock, _ := rm.lockById.LoadOrStore(roomId, &sync.RWMutex{})

	return lock.(*sync.RWMutex)
}
