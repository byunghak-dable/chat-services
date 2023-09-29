package db

import (
	"context"

	"github.com/go-redis/redis/v9"
	"github.com/widcraft/user-service/pkg/logger"
)

type Redis struct {
	logger logger.Logger
	*redis.Client
	ctx context.Context
}

func NewRedis(logger logger.Logger, address, password string, db int) (*Redis, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}
	return &Redis{
		logger: logger,
		Client: rdb,
		ctx:    ctx,
	}, nil
}

func (db *Redis) Close() error {
	return db.Conn().Close()
}
