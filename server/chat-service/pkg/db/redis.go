package db

import (
	"context"

	"github.com/go-redis/redis/v9"
	"github.com/widcraft/chat-service/pkg/logger"
)

type Redis struct {
	logger logger.Logger
	*redis.Client
	ctx context.Context
}

func NewRedis(logger logger.Logger, address, password string, db int) (*Redis, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}
	return &Redis{
		logger: logger,
		Client: client,
		ctx:    ctx,
	}, nil
}

func (db *Redis) Close() error {
	return db.Conn().Close()
}
