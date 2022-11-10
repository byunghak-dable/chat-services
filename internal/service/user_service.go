package service

import "log"

type UserService struct {
	logger *log.Logger
}

func New(logger *log.Logger) *UserService {
	return &UserService{
		logger: logger,
	}
}

func (a UserService) Register() {

}

func (a UserService) GetUserByIdx() {

}
