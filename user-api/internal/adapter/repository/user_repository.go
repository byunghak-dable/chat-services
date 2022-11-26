package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/weed082/user-api/internal/adapter/repository/mysql"
	"github.com/weed082/user-api/internal/domain/dto"
	"github.com/weed082/user-api/internal/domain/entity"
)

type UserRepo struct {
	logger log.FieldLogger
	db     *mysql.Mysql
}

func NewUserRepo(logger log.FieldLogger, db *mysql.Mysql) *UserRepo {
	return &UserRepo{
		logger: logger,
		db:     db,
	}
}

func (repo UserRepo) Register(registerDto dto.RegisterReqDto) error {
	return nil
}

func (repo UserRepo) Signin(signinDto dto.SigninReqDto) (*entity.User, error) {
	return nil, nil
}
