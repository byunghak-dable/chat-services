package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/weed082/chat-server/internal/adapter/driven/repository/mysql"
	"github.com/weed082/chat-server/internal/domain/dto"
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

func (repo UserRepo) Register(registerDto dto.RegisterDto) error {
	return nil
}
