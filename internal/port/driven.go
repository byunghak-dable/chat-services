package port

import "github.com/weed082/chat-server/internal/domain/dto"

type UserRepo interface {
	Register(registerDto dto.RegisterDto) error
}
