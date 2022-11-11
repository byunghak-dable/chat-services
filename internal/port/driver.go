package port

import "github.com/weed082/chat-server/internal/domain/dto"

type UserApp interface {
	Register(registerDto dto.RegisterDto) error
}
