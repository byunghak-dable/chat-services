package main

import (
	"net"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/adapter/grpc"
	"github.com/widcraft/chat-service/internal/adapter/repository"
	"github.com/widcraft/chat-service/internal/adapter/repository/redis"
	"github.com/widcraft/chat-service/internal/adapter/rest"
	chatapp "github.com/widcraft/chat-service/internal/application/chat"
)

type Closable interface {
	Close() error
}

// env
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
}

func main() {
	logger := log.New()
	redisDb, err := redis.New(logger, net.JoinHostPort(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASSWORD"), 0)
	defer shutdown(logger, redisDb)

	if err != nil {
		logger.Error(err)
		return
	}

	chatRepo := repository.NewChatRepository(logger, redisDb)
	chatApp := chatapp.New(logger, chatRepo)
	restServer := rest.New(logger, chatApp)
	grpcServer := grpc.New(logger, chatApp)
	defer shutdown(logger, restServer, grpcServer)

	go restServer.Run(os.Getenv("REST_PORT"))
	go grpcServer.Run(os.Getenv("GRPC_PORT"))

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}

func shutdown(logger log.FieldLogger, closables ...Closable) {
	for _, target := range closables {
		if reflect.ValueOf(target).IsNil() {
			continue
		}
		if err := target.Close(); err != nil {
			logger.Errorf("%s closing failed: %s", reflect.TypeOf(target), err)
		}
	}
}
