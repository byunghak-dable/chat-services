package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v9"
)

type Redis struct {
	logger *log.Logger
	*redis.Client
	ctx context.Context
}

func New(logger *log.Logger, address, password string, db int) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db, // use default DB
	})

	return &Redis{
		logger: logger,
		Client: rdb,
		ctx:    context.Background(),
	}
}
