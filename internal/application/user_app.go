package application

import (
	"log"

	"github.com/weed082/chat-server/internal/port"
)

type UserApp struct {
	logger *log.Logger
	repo   port.UserRepo
}

func NewUserApp(logger *log.Logger, repo port.UserRepo) *UserApp {
	return &UserApp{
		logger: logger,
		repo:   repo,
	}
}

func (a UserApp) Register() {

}

func (a UserApp) GetUserByIdx() {

}
