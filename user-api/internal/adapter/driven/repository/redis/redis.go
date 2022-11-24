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

func (db *Redis) Close() {
	if err := db.Conn().Close(); err != nil {
		db.logger.Errorf("close redis failure: %v", err)
	}
	db.logger.Infoln("redis successfully closed")
}
