package main

import (
	"context"
	"messenger-service/internal/adapter/driven"
	"messenger-service/internal/adapter/driven/config"
	"messenger-service/internal/adapter/driven/persistence"
	"messenger-service/internal/adapter/driven/persistence/db"
	"messenger-service/internal/adapter/driving"
	"messenger-service/internal/adapter/driving/grpc"
	"messenger-service/internal/adapter/driving/rest"
	"messenger-service/internal/application/dto"
	"messenger-service/internal/application/mapper"
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

	configStore := load(config.NewStore())
	mongoDb := loadClosable(db.NewMongoDb(configStore))
	kafkaProducer := loadClosable(driven.NewKafkaProducer[dto.Message](configStore))

	messageRepository := persistence.NewMessageRepository(logger, mongoDb)
	messageStore := service.NewMessageStore(messageRepository, mapper.NewMessage())
	messenger := service.NewMessenger(kafkaProducer, messageStore, service.NewRoomManager())

	kafkaConsumer := loadClosable(driving.NewKafkaConsumer[dto.Message](configStore, logger, messenger))
	restApp := loadClosable(rest.New(configStore, logger, messenger, messageStore), nil)
	grpcApp := loadClosable(grpc.New(configStore, logger, messenger), nil)

	run(kafkaConsumer, restApp, grpcApp)
}

func loadClosable[T Closable](target T, err error) T {
	closableList = append(closableList, target)

	return load(target, err)
}

func load[T any](target T, err error) T {
	if err != nil {
		logger.Errorf("load failed: %v", err)
		panic(err)
	}

	return target
}

func run(runnables ...Runnable) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer waitTermination(ctx)

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
			logger.Errorf("%s close failed: %s", reflect.TypeOf(closable), err)
			continue
		}
		logger.Infof("%s successfully closed", reflect.TypeOf(closable))
	}
}

func waitTermination(ctx context.Context) {
	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		logger.Infoln("Received context done")
	case s := <-terminationChan:
		logger.Infof("Received signal: %s", s)
	}
}
