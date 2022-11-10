package repository

import (
	"log"

	"github.com/weed082/chat-server/internal/adapter/driven/repository/mysql"
)

type UserRepo struct {
	logger *log.Logger
	db     *mysql.Mysql
}

func NewUserRepo(logger *log.Logger, db *mysql.Mysql) *UserRepo {
	return &UserRepo{
		logger: logger,
		db:     db,
	}
}
