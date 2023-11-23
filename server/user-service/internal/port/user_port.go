package port

import (
	"github.com/widcraft/user-service/internal/domain/dto"
	"github.com/widcraft/user-service/internal/domain/entity"
)

type UserApp interface {
	Register(*dto.RegisterReqDto) error
}

type UserRepository interface {
	Register(*entity.User) error
}
