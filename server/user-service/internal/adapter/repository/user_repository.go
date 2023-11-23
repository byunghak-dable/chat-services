package repository

import (
	"github.com/widcraft/user-service/internal/domain/entity"
	"github.com/widcraft/user-service/pkg/db"
	"github.com/widcraft/user-service/pkg/logger"
)

type UserRepo struct {
	logger logger.Logger
	db     *db.Mysql
}

func NewUserRepo(logger logger.Logger, db *db.Mysql) *UserRepo {
	return &UserRepo{
		logger: logger,
		db:     db,
	}
}

func (repo *UserRepo) Register(user *entity.User) error {
	return nil
}

func (repo *UserRepo) GoogleSignin(token string) (*entity.User, error) {
	return nil, nil
}
