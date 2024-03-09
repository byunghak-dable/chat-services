package service

import (
	"errors"
	"fmt"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driving"
	"sync"
)

type RoomManager struct {
	lockMap *sync.Map
	roomMap map[uint][]driving.MessengerClient
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		lockMap: new(sync.Map),
		roomMap: make(map[uint][]driving.MessengerClient),
	}
}

func (rm *RoomManager) Join(client driving.MessengerClient) error {
	roomIdx := client.GetRoomIdx()

	return rm.withLock(roomIdx, func() error {
		rm.roomMap[roomIdx] = append(rm.roomMap[roomIdx], client)
		return nil
	})
}

func (rm *RoomManager) Leave(client driving.MessengerClient) error {
	roomIdx := client.GetRoomIdx()

	return rm.withLock(roomIdx, func() error {
		participants := rm.roomMap[roomIdx]

		for i, participant := range rm.roomMap[roomIdx] {
			if client != participant {
				continue
			}

			rm.roomMap[roomIdx] = append(participants[:i], participants[i+1:]...)

			if len(rm.roomMap[roomIdx]) == 0 {
				delete(rm.roomMap, roomIdx)
				rm.lockMap.Delete(roomIdx)
			}

			return nil
		}

		return errors.New("no matching client to leave")
	})
}

func (rm *RoomManager) Broadcast(message *dto.Message) error {
	roomIdx := message.RoomIdx

	return rm.withRLock(roomIdx, func() error {
		var errs []error
		room := rm.roomMap[roomIdx]
		errChan := make(chan error)
		defer close(errChan)

		for _, participant := range room {
			go func(participant driving.MessengerClient) {
				errChan <- participant.SendMessage(message)
			}(participant)
		}

		for range room {
			if err := <-errChan; err != nil {
				errs = append(errs, err)
			}
		}

		if len(errs) > 0 {
			return fmt.Errorf("broadcast failed: %v", errs)
		}

		return nil
	})
}

func (rm *RoomManager) withLock(roomIdx uint, action func() error) error {
	lock := rm.getLock(roomIdx)
	lock.Lock()
	defer lock.Unlock()

	return action()
}

func (rm *RoomManager) withRLock(roomIdx uint, action func() error) error {
	lock := rm.getLock(roomIdx)
	lock.RLock()
	defer lock.RUnlock()

	return action()
}

func (rm *RoomManager) getLock(roomIdx uint) *sync.RWMutex {
	mutex, _ := rm.lockMap.LoadOrStore(roomIdx, &sync.RWMutex{})

	return mutex.(*sync.RWMutex)
}
