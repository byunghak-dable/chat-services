package main

import (
	"context"
	"fmt"
	"messenger-service/internal/adapter/driven"
	"messenger-service/internal/adapter/driven/persistence"
	"messenger-service/internal/adapter/driven/persistence/db"
	"messenger-service/internal/adapter/driving"
	"messenger-service/internal/adapter/driving/grpc"
	"messenger-service/internal/adapter/driving/rest"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/application/service"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type Runnable interface {
	Run() error
}

type Closable interface {
	Close() error
}

var logger = log.New()
var closableList []Closable

func main() {
	defer quit()

	configStore, err := driven.NewConfigStore()

	if err != nil {
		logger.Error(err)
		return
	}

	mongoDb := loadClosable(db.NewMongoDb(configStore))
	kafkaProducer := loadClosable(driven.NewKafkaProducer[dto.Message](configStore))

	messageRepository := persistence.NewMessageRepository(mongoDb)
	messageStore := service.NewMessageStore(messageRepository)
	messenger := service.NewMessenger(kafkaProducer, messageStore, service.NewRoomManager())

	kafkaConsumer := loadClosable(driving.NewKafkaConsumer(configStore, logger, messenger))
	restApp := loadClosable(rest.New(configStore, logger, messenger), nil)
	grpcApp := loadClosable(grpc.New(configStore, logger, messenger), nil)

	run(kafkaConsumer, restApp, grpcApp)
}

func loadClosable[T Closable](target T, err error) T {
	closableList = append(closableList, target)

	if err != nil {
		panic(err)
	}

	return target
}

func run(runnables ...Runnable) {
	ctx, cancel := context.WithCancel(context.Background())

	defer waitTermination(ctx)
	defer cancel()

	for _, runnable := range runnables {
		go func(runnable Runnable) {
			if err := runnable.Run(); err != nil {
				logger.Errorf("%s failed: %s", reflect.TypeOf(runnable), err)
				cancel()
			}
		}(runnable)
	}
}

func quit() {
	for _, closable := range closableList {
		if reflect.ValueOf(closable).IsNil() {
			continue
		}

		if err := closable.Close(); err != nil {
			logger.Errorf("%s exiting failed: %s", reflect.TypeOf(closable), err)
			return
		}
		logger.Infof("%s successfully closed", reflect.TypeOf(closable))
	}
}

func waitTermination(ctx context.Context) {
	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		fmt.Println("Received context done")
	case s := <-terminationChan:
		fmt.Println("Received signal:", s)
	}
}
