package application

import (
	"github.com/widcraft/user-service/internal/domain/dto"
	"github.com/widcraft/user-service/internal/domain/entity"
	"github.com/widcraft/user-service/internal/port"
	"github.com/widcraft/user-service/pkg/logger"
)

type UserApp struct {
	logger     logger.Logger
	repository port.UserRepository
}

func NewUserApp(logger logger.Logger, repository port.UserRepository) *UserApp {
	return &UserApp{
		logger:     logger,
		repository: repository,
	}
}

func (app *UserApp) Register(registerDto *dto.RegisterReqDto) error {
	return app.repository.Register(&entity.User{
		Email:    registerDto.Email,
		Name:     registerDto.Name,
		ImageUrl: registerDto.ImageUrl,
		Token:    registerDto.Token,
	})
}
