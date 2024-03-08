package service

import (
	"errors"
	"fmt"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driving"
	"sync"
)

type RoomService struct {
	lockMap *sync.Map
	roomMap map[uint][]driving.MessengerClient
}

func NewRoomService() *RoomService {
	return &RoomService{
		lockMap: new(sync.Map),
		roomMap: make(map[uint][]driving.MessengerClient),
	}
}

func (service *RoomService) Join(client driving.MessengerClient) error {
	roomIdx := client.GetRoomIdx()

	return service.withLock(roomIdx, func() error {
		service.roomMap[roomIdx] = append(service.roomMap[roomIdx], client)
		return nil
	})
}

func (service *RoomService) Leave(client driving.MessengerClient) error {
	roomIdx := client.GetRoomIdx()

	return service.withLock(roomIdx, func() error {
		participants := service.roomMap[roomIdx]

		for i, participant := range service.roomMap[roomIdx] {
			if client != participant {
				continue
			}

			service.roomMap[roomIdx] = append(participants[:i], participants[i+1:]...)

			if len(service.roomMap[roomIdx]) == 0 {
				delete(service.roomMap, roomIdx)
				service.lockMap.Delete(roomIdx)
			}

			return nil
		}

		return errors.New("no matching client to leave")
	})
}

func (service *RoomService) Broadcast(message *dto.MessageDto) error {
	roomIdx := message.RoomIdx

	return service.withRLock(roomIdx, func() error {
		var errs []error

		for _, participant := range service.roomMap[roomIdx] {
			if err := participant.SendMessage(message); err != nil {
				errs = append(errs, err)
			}
		}

		if len(errs) > 0 {
			return fmt.Errorf("bradcast error: %v", errs)
		}

		return nil
	})
}

func (service *RoomService) withLock(roomIdx uint, action func() error) error {
	lock := service.getLock(roomIdx)
	lock.Lock()
	defer lock.Unlock()

	return action()
}

func (service *RoomService) withRLock(roomIdx uint, action func() error) error {
	lock := service.getLock(roomIdx)
	lock.RLock()
	defer lock.RUnlock()

	return action()
}

func (service *RoomService) getLock(roomIdx uint) *sync.RWMutex {
	mutex, _ := service.lockMap.LoadOrStore(roomIdx, &sync.RWMutex{})

	return mutex.(*sync.RWMutex)
}
