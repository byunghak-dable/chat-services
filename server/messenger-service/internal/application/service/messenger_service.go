package service

import (
	"fmt"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
	"sync"
)

type MessengerService struct {
	rooms    map[uint][]driving.MessengerClientPort
	mutex    *sync.RWMutex
	logger   driven.LoggerPort
	producer driven.MessageProducerPort
}

func NewMessengerService(logger driven.LoggerPort, producer driven.MessageProducerPort) *MessengerService {
	return &MessengerService{
		rooms:    make(map[uint][]driving.MessengerClientPort),
		mutex:    new(sync.RWMutex),
		logger:   logger,
		producer: producer,
	}
}

func (service *MessengerService) Join(client driving.MessengerClientPort) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	roomIdx := client.GetRoomIdx()
	room, ok := service.rooms[roomIdx]

	if ok {
		service.rooms[roomIdx] = append(room, client)
		service.logger.Infof("room id: %d, count: %d", roomIdx, len(room))
		return
	}

	service.rooms[roomIdx] = []driving.MessengerClientPort{client}
	service.logger.Infof("room id: %d added", roomIdx, len(room))
}

func (service *MessengerService) Leave(client driving.MessengerClientPort) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	roomIdx := client.GetRoomIdx()
	room, ok := service.rooms[roomIdx]

	if !ok {
		service.logger.Error("no existing chat room roomIdx")
		return
	}

	for i, participant := range room {
		if client == participant {
			service.rooms[roomIdx] = append(room[:i], room[i+1:]...)
			return
		}
	}
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

	// TODO: need better error handling
	if len(sendErrors) > 0 {
		return fmt.Errorf("send message errors, %v", sendErrors)
	}

	return nil
}

func (service *MessengerService) ProduceMessage(message *dto.MessageDto) error {
	return service.producer.ProduceMessage(message)
}
