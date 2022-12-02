package rest

import (
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/port"
)

type Rest struct {
	logger  *log.Logger
	chatApp port.ChatApp
}

func NewRest(logger *log.Logger, chatApp port.ChatApp) *Rest {
	return &Rest{
		logger:  logger,
		chatApp: chatApp,
	}
}

func (rest *Rest) Close() {
	// need to close websocket gracefully
}
