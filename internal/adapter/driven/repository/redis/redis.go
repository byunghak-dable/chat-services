package redis

import "log"

type Redis struct {
	logger *log.Logger
}

func New(logger *log.Logger) *Redis {
	return &Redis{
		logger: logger,
	}
}
