package main

import (
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/internal/adapter/primary/grpc"
	"github.com/widcraft/chat-service/internal/adapter/primary/rest"
	"github.com/widcraft/chat-service/internal/adapter/secondary/persistence/db"
	"github.com/widcraft/chat-service/internal/adapter/secondary/persistence/repository"
	"github.com/widcraft/chat-service/internal/application"
	"github.com/widcraft/chat-service/internal/application/message"
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
	mongoDb, err := db.NewMongoDb(db.MongoDbConfig{
		User:     os.Getenv("MONGODB_USER"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		Host:     os.Getenv("MONGODB_HOST"),
		Port:     os.Getenv("MONGODB_PORT"),
	})
	if err != nil {
		logger.Error(err)
		return
	}

	messageServiceFacade := application.NewChatService(
		logger,
		message.NewMessageService(logger, repository.NewMessageRepository(logger, mongoDb)),
		message.NewMessengerService(logger),
	)

	restApp := rest.New(logger, messageServiceFacade)
	grpcApp := grpc.New(logger, messageServiceFacade)

	defer shutdown(restApp, grpcApp)

	go restApp.Run(os.Getenv("REST_PORT"))
	go grpcApp.Run(os.Getenv("GRPC_PORT"))

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
