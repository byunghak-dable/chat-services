package db

import (
	"context"
	"net"

	"github.com/go-redis/redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Db       int
}

type Redis struct {
	ctx context.Context
	*redis.Client
}

func NewRedis(config RedisConfig) (*Redis, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(config.Host, config.Port),
		Password: config.Password,
		DB:       config.Db,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &Redis{
		Client: client,
		ctx:    ctx,
	}, nil
}

func (db *Redis) Close() error {
	return db.Conn().Close()
}
