package service

import "message-service/internal/port/driven"

type MessageService struct {
	repository driven.MessageRepository
}

func NewMessageService(messageRepository driven.MessageRepository) *MessageService {
	return &MessageService{
		repository: messageRepository,
	}
}
