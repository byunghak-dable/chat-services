package application

import (
	log "github.com/sirupsen/logrus"
	"github.com/weed082/chat-server/internal/domain/dto"
	"github.com/weed082/chat-server/internal/port"
)

type UserApp struct {
	logger log.FieldLogger
	repo   port.UserRepo
}

func NewUserApp(logger log.FieldLogger, repo port.UserRepo) *UserApp {
	return &UserApp{
		logger: logger,
		repo:   repo,
	}
}

func (app UserApp) Register(registerDto dto.RegisterDto) error {
	return nil
}
