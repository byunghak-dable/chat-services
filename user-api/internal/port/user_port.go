package port

import "github.com/weed082/user-api/internal/domain/dto"

type UserApp interface {
	Register(dto.RegisterDto) error
	Signin(dto.SigninDto) error
}

type UserRepo interface {
}
