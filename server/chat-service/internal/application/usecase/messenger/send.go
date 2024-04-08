package messenger

import (
	"chat-service/internal/application/abstraction"
	"chat-service/internal/application/dto"
	"chat-service/internal/application/mapper"
	"chat-service/internal/port/driven"
)

type SendUseCase struct {
	logger      driven.Logger
	broker      driven.MessageBroker
	repository  driven.MessageRepository
	roomManager abstraction.RoomManager
	mapper      *mapper.Message
}

func NewSendUseCase(logger driven.Logger, broker driven.MessageBroker, repository driven.MessageRepository, roomManager abstraction.RoomManager, mapper *mapper.Message) *SendUseCase {
	useCase := &SendUseCase{logger, broker, repository, roomManager, mapper}

	broker.Subscribe(useCase)

	return useCase
}

func (u *SendUseCase) Handle(message dto.Message) error {
	entity := u.mapper.ToEntity(message)

	if err := u.repository.Create(&entity); err != nil {
		return err
	}

	return u.broker.Publish(u.mapper.ToDto(entity))
}

func (u *SendUseCase) OnReceive(message dto.Message) {
	if err := u.roomManager.Broadcast(message); err != nil {
		u.logger.Errorf("[SEND_USE_CASE] broadcast failed: %v", err)
	}
}
