package main

import (
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/adapter/grpc"
	"github.com/widcraft/chat-service/internal/adapter/repository"
	"github.com/widcraft/chat-service/internal/adapter/rest"
	chatapp "github.com/widcraft/chat-service/internal/application/chat"
	"github.com/widcraft/chat-service/pkg/db"
)

var logger = log.New()

type Closable interface {
	Close() error
}

// env
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	redisConfig := db.RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Db:       0,
	}
	redisDb, err := db.NewRedis(redisConfig)

	defer shutdown(redisDb)

	if err != nil {
		logger.Error(err)
		return
	}

	chatRepo := repository.NewChatRepository(logger, redisDb)
	chatApp := chatapp.New(logger, chatRepo)
	restServer := rest.New(logger, chatApp)
	grpcServer := grpc.New(logger, chatApp)

	defer shutdown(restServer, grpcServer)

	go restServer.Run(os.Getenv("REST_PORT"))
	go grpcServer.Run(os.Getenv("GRPC_PORT"))

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}

func shutdown(closables ...Closable) {
	for _, closable := range closables {
		if reflect.ValueOf(closable).IsNil() {
			continue
		}

		err := closable.Close()
		if err != nil {
			logger.Errorf("%s closing failed: %s", reflect.TypeOf(closable), err)
		}
	}
}
