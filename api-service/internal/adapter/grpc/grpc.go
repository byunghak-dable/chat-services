package grpc

import (
	log "github.com/sirupsen/logrus"
)

type Grpc struct {
	logger log.FieldLogger
}

func New(logger log.FieldLogger) *Grpc {
	return &Grpc{
		logger: logger,
	}
}
