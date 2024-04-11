package main

import (
	"chat-service/internal/adapter/driven/config"
	"chat-service/internal/adapter/driven/messaging"
	"chat-service/internal/adapter/driven/persistence/db"
	"chat-service/internal/adapter/driven/persistence/repository"
	"chat-service/internal/adapter/driver/rest"
	hmessage "chat-service/internal/adapter/driver/rest/message"
	hmessenger "chat-service/internal/adapter/driver/rest/messenger"
	"chat-service/internal/application/mapper"
	"chat-service/internal/application/usecase/message"
	"chat-service/internal/application/usecase/messenger"
	"chat-service/internal/domain/service"
	"context"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type Runnable interface {
	Run(ctx context.Context) error
}

type Closable interface {
	Close() error
}

var (
	logger    = log.New()
	closables []Closable
	runnables []Runnable
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	defer func() {
		exit()
		cancel()
		wg.Wait()
	}()

	initialize()
	defer waitExitSignal(ctx)

	for _, runnable := range runnables {
		wg.Add(1)

		go func(runnable Runnable) {
			defer wg.Done()

			runnableType := reflect.TypeOf(runnable)

			if err := runnable.Run(ctx); err != nil {
				cancel()
				logger.Errorf("[MAIN] run %s failed %s", runnableType, err)
				return
			}

			logger.Infof("[MAIN] running %s exited", runnableType)
		}(runnable)
	}
}

func initialize() {
	configStore := load(config.New())
	mongoDb := load(db.NewMongoDb(configStore))
	restServer := load(rest.New(configStore), nil)

	// adapter
	messageBroker := load(messaging.NewMessageBroker(configStore, logger))
	messageRepository := repository.NewMessage(logger, mongoDb)

	// mapper
	messageMapper := mapper.NewMessage()

	// domain
	roomManager := service.NewRoomManager()

	// use case
	messageRead := message.NewReadMultiUseCase(messageRepository, messageMapper)

	messengerJoin := messenger.NewJoinUseCase(roomManager)
	messengerLeave := messenger.NewLeaveUseCase(roomManager)
	messengerSend := messenger.NewSendUseCase(logger, messageBroker, messageRepository, roomManager, messageMapper)

	restServer.Register(
		hmessage.NewHandler(logger, messageRead),
		hmessenger.NewHandler(logger, messengerJoin, messengerLeave, messengerSend),
	)
}

func load[T any](target T, err error) T {
	if err != nil {
		logger.Errorf("[MAIN] load failed: %v", err)
		panic(err)
	}

	targetInterface := reflect.ValueOf(target).Interface()

	if closable, ok := targetInterface.(Closable); ok {
		closables = append(closables, closable)
	}

	if runnable, ok := targetInterface.(Runnable); ok {
		runnables = append(runnables, runnable)
	}

	return target
}

func waitExitSignal(ctx context.Context) {
	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		logger.Infoln("[MAIN] received context done")
	case s := <-terminationChan:
		logger.Infof("[MAIN] received signal: %s", s)
	}
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
