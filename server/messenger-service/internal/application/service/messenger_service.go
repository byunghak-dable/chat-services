package service

import (
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
)

type MessengerService struct {
	logger      driven.LoggerPort
	producer    driven.MessageProducerPort
	roomService *RoomService
}

func NewMessengerService(logger driven.LoggerPort, producer driven.MessageProducerPort, roomService *RoomService) *MessengerService {
	return &MessengerService{
		logger:      logger,
		producer:    producer,
		roomService: roomService,
	}
}

func (service *MessengerService) Join(client driving.MessengerClientPort) error {
	return service.roomService.Join(client)
}

func (service *MessengerService) Leave(client driving.MessengerClientPort) error {
	return service.roomService.Leave(client)
}

func (service *MessengerService) Broadcast(message *dto.MessageDto) error {
	return service.roomService.Broadcast(message)
}

func (service *MessengerService) SendMessage(message *dto.MessageDto) error {
	return service.producer.Produce(message)
}
