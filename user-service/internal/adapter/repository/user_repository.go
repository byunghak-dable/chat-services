package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/user-service/internal/adapter/repository/mysql"
	"github.com/widcraft/user-service/internal/domain/entity"
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
	result := repo.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo UserRepo) GoogleSignin(token string) (*entity.User, error) {
	return nil, nil
}
