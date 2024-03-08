package main

import (
	"context"
	"fmt"
	"messenger-service/internal/adapter/driven"
	"messenger-service/internal/adapter/driving"
	"messenger-service/internal/adapter/driving/grpc"
	"messenger-service/internal/adapter/driving/rest"
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

func main() {
	configStore, err := driven.NewConfigStore()

	if err != nil {
		logger.Error(err)
		return
	}

	kafkaProducer, producerErr := driven.NewMessageProducer(configStore)

	defer quit(kafkaProducer)

	if producerErr != nil {
		logger.Error(producerErr)
		return
	}

	messengerService := service.NewMessengerService(logger, kafkaProducer)

	kafkaConsumer, consumerErr := driving.NewMessageBroadcaster(configStore, logger, messengerService)
	restApp := rest.New(configStore, logger, messengerService)
	grpcApp := grpc.New(configStore, logger, messengerService)

	defer quit(restApp, grpcApp, kafkaProducer)

	if consumerErr != nil {
		logger.Error(consumerErr)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	run(cancel, kafkaConsumer, restApp, grpcApp)
	handleTermination(ctx)
}

func run(cancel context.CancelFunc, runnables ...Runnable) {
	for _, runnable := range runnables {
		go func(runnable Runnable) {
			if err := runnable.Run(); err != nil {
				logger.Errorf("%s failed: %s", reflect.TypeOf(runnable), err)
				cancel()
			}
		}(runnable)
	}
}

func quit(closables ...Closable) {
	for _, closable := range closables {
		if reflect.ValueOf(closable).IsNil() {
			continue
		}

		err := closable.Close()
		if err != nil {
			logger.Errorf("%s exiting failed: %s", reflect.TypeOf(closable), err)
		}
	}
}

func handleTermination(ctx context.Context) {
	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		fmt.Println("Received context done")
	case s := <-terminationChan:
		fmt.Println("Received signal:", s)
	}
}
