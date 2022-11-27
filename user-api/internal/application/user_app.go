package application

import (
	log "github.com/sirupsen/logrus"
	"github.com/weed082/user-api/internal/domain/dto"
	"github.com/weed082/user-api/internal/domain/entity"
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

func (app UserApp) Register(registerDto *dto.RegisterReqDto) error {
	return app.repository.Register(&entity.User{
		Email:    registerDto.Email,
		Name:     registerDto.Name,
		ImageUrl: registerDto.ImageUrl,
		Token:    registerDto.Token,
	})
}

func (app UserApp) GoogleSignin(signinDto *dto.GoogleSigninReqDto) (*dto.GoogleSigninResDto, error) {
	return nil, nil
}
