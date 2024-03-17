package mapper

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/entity"
)

type Message struct {
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) ToEntity(dto dto.Message) *entity.Message {
	return &entity.Message{
		Id:        dto.Id,
		Message:   dto.Message,
		RoomId:    dto.RoomId,
		UserId:    dto.UserId,
		UpdatedAt: dto.UpdatedAt,
	}
}

func (m *Message) ToDto(entity *entity.Message) dto.Message {
	return dto.Message{
		Id:        entity.Id,
		Message:   entity.Message,
		RoomId:    entity.RoomId,
		UserId:    entity.UserId,
		UpdatedAt: entity.UpdatedAt,
	}
}

func (m *Message) ToDtoList(entities []*entity.Message) []dto.Message {
	var dtos []dto.Message

	for _, message := range entities {
		dtos = append(dtos, m.ToDto(message))
	}

	return dtos
}
