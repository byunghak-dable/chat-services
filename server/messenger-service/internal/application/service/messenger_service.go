package service

import (
	"messenger-service/internal/application/dto"
	"messenger-service/internal/port/driven"
	"messenger-service/internal/port/driving"
)

type MessengerService struct {
	logger      driven.Logger
	producer    driven.MessageProducer
	roomService *RoomService
}

func NewMessengerService(logger driven.Logger, producer driven.MessageProducer) *MessengerService {
	return &MessengerService{
		logger:      logger,
		producer:    producer,
		roomService: NewRoomService(),
	}
}

func (service *MessengerService) Join(client driving.MessengerClient) error {
	return service.roomService.Join(client)
}

func (service *MessengerService) Leave(client driving.MessengerClient) error {
	return service.roomService.Leave(client)
}

func (service *MessengerService) Broadcast(message *dto.MessageDto) error {
	return service.roomService.Broadcast(message)
}

func (service *MessengerService) SendMessage(message *dto.MessageDto) error {
	return service.producer.Produce(message)
}
