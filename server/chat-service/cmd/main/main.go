package main

import (
	"chat-service/internal/adapter/driven/config"
	"chat-service/internal/adapter/driven/messaging"
	"chat-service/internal/adapter/driven/persistence"
	"chat-service/internal/adapter/driven/persistence/db"
	"chat-service/internal/adapter/driver/rest"
	"chat-service/internal/application/mapper"
	"chat-service/internal/application/service"
	"context"
	"io"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type Runnable interface {
	Run() error
}

var logger = log.New()
var closables []io.Closer

func main() {
	defer exit()

	configStore, err := config.New()
	if err != nil {
		logger.Errorf("[MAIN] load config failed: %v", err)
		panic(err)
	}

	mongoDb := load(db.NewMongoDb(configStore))
	messageBroker := load(messaging.NewMessageBroker(configStore, logger))

	messageRepo := persistence.NewMessageRepository(logger, mongoDb)
	message := service.NewMessage(messageRepo, mapper.NewMessage())
	messenger := service.NewMessenger(logger, messageBroker, message)

	restApp := load(rest.New(configStore, logger, messenger, message), nil)

	run(messageBroker, restApp)
}

func load[T io.Closer](target T, err error) T {
	closables = append(closables, target)

	if err != nil {
		logger.Errorf("[MAIN] load failed: %v", err)
		panic(err)
	}

	return target
}

func exit() {
	for _, closable := range closables {
		if reflect.ValueOf(closable).IsNil() {
			continue
		}

		if err := closable.Close(); err != nil {
			logger.Errorf("[MAIN] %s exit failed: %s", reflect.TypeOf(closable), err)
			continue
		}

		logger.Infof("[MAIN] %s successfully closed", reflect.TypeOf(closable))
	}
}

func run(runnables ...Runnable) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer waitTermination(ctx)

	for _, runnable := range runnables {
		go func(runnable Runnable) {
			if err := runnable.Run(); err != nil {
				logger.Errorf("[MAIN] %s run failed: %s", reflect.TypeOf(runnable), err)
				cancel()
			}
		}(runnable)
	}
}

func waitTermination(ctx context.Context) {
	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		logger.Infoln("[MAIN] received context done")
	case s := <-terminationChan:
		logger.Infof("[MAIN] received signal: %s", s)
	}
}
