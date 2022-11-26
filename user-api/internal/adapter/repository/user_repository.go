package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/weed082/user-api/internal/adapter/repository/mysql"
	"github.com/weed082/user-api/internal/domain/entity"
)

type UserRepo struct {
	logger log.FieldLogger
	db     *mysql.Mysql
}

func NewUserRepo(logger log.FieldLogger, db *mysql.Mysql) *UserRepo {
	db.AutoMigrate(&entity.User{})

	return &UserRepo{
		logger: logger,
		db:     db,
	}
}

func (repo UserRepo) Register(user *entity.User) error {
	return nil
}

func (repo UserRepo) GoogleSignin(token string) (*entity.User, error) {
	return nil, nil
}
