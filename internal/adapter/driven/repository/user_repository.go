package repository

import "log"

type UserRepository struct {
	logger *log.Logger
}

func New(logger *log.Logger) *UserRepository {
	return &UserRepository{
		logger: logger,
	}
}
