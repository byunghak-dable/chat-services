package application

import (
	log "github.com/sirupsen/logrus"
	"github.com/weed082/user-api/internal/domain/dto"
	"github.com/weed082/user-api/internal/port"
)

type UserApp struct {
	logger     log.FieldLogger
	repository port.UserRepository
}

func NewUserApp(logger log.FieldLogger, repository port.UserRepository) *UserApp {
	return &UserApp{
		logger:     logger,
		repository: repository,
	}
}

func (app UserApp) Register(registerDto dto.RegisterReqDto) error {
	return nil
}

func (app UserApp) Signin(signinDto dto.SigninReqDto) (*dto.SigninResDto, error) {
	return nil, nil
}
