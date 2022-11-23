package port

import "github.com/weed082/chat-server/internal/domain/dto"

type UserApp interface {
	Register(dto.RegisterDto) error
	Signin(dto.SigninDto) error
}
