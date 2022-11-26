package port

import "github.com/weed082/user-api/internal/domain/dto"

type UserApp interface {
	Register(dto.RegisterReqDto) error
	Signin(dto.SigninReqDto) (*dto.SigninResDto, error)
}

type UserRepository interface {
	Register(dto.RegisterReqDto) error
	Signin(dto.SigninReqDto) (*dto.SigninResDto, error)
}
