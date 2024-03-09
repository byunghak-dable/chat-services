package mapper

import (
	"messenger-service/internal/application/dto"
	"messenger-service/internal/domain/entity"
)

type Message struct {
}

func NewMessage() *Message {
	return &Message{}
}

func (mm *Message) ToEntity(dto *dto.Message) *entity.Message {
	return &entity.Message{
		Message: dto.Message,
		RoomIdx: dto.RoomIdx,
		UserIdx: dto.UserIdx,
	}

}

func (mm *Message) ToDto(entity *entity.Message) *dto.Message {
	return &dto.Message{
		Message: entity.Message,
		RoomIdx: entity.RoomIdx,
		UserIdx: entity.UserIdx,
	}
}

func (mm *Message) ToDtoList(entities []*entity.Message) []*dto.Message {
	var dtoList []*dto.Message

	for _, message := range entities {
		dtoList = append(dtoList, mm.ToDto(message))
	}

	return dtoList
}
