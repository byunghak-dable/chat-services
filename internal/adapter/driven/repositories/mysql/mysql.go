package mysql

import "log"

type Mysql struct {
	logger *log.Logger
}

func New(logger *log.Logger) *Mysql {
	return &Mysql{
		logger: logger,
	}
}
