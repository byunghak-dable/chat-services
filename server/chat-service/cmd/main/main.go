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
	Run(ctx context.Context, wg *sync.WaitGroup) error
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

	defer terminate(&wg, cancel)

	run(&wg, ctx)
	waitTermination(ctx)
}

func run(wg *sync.WaitGroup, ctx context.Context) {
	initialize()

	for _, runnable := range runnables {
		wg.Add(1)

		go func(runnable Runnable) {
			if err := runnable.Run(ctx, wg); err != nil {
				logger.Errorf("[MAIN] %s run failed: %s", reflect.TypeOf(runnable), err)
				panic(err)
			}
		}(runnable)
	}
}

func initialize() {
	configStore := load(config.New())
	mongoDb := load(db.NewMongoDb(configStore))
	restServer := load(rest.New(configStore, logger), nil)

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

func terminate(wg *sync.WaitGroup, cancel context.CancelFunc) {
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

	cancel()
	wg.Wait()
}
