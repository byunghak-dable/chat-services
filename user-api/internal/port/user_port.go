package port

import (
	"github.com/weed082/user-api/internal/domain/dto"
	"github.com/weed082/user-api/internal/domain/entity"
)

type UserApp interface {
	Register(*dto.RegisterReqDto) error
	GoogleSignin(*dto.GoogleSigninReqDto) (*dto.GoogleSigninResDto, error)
}

type UserRepository interface {
	Register(*entity.User) error
	GoogleSignin(string) (*entity.User, error)
}
