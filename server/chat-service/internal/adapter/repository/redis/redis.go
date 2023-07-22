package redis

import (
	"context"

	"github.com/go-redis/redis/v9"
	log "github.com/sirupsen/logrus"
)

type Redis struct {
	logger log.FieldLogger
	*redis.Client
	ctx context.Context
}

func New(logger log.FieldLogger, address, password string, db int) (*Redis, error) {
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
