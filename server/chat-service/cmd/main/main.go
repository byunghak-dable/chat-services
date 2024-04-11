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
)

func main() {
	configStore := load(config.New())
	mongoDb := loadClosable(db.NewMongoDb(configStore))
	restServer := loadClosable(rest.New(configStore, logger), nil)

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

	run(messageBroker, restServer)
}

func loadClosable[T Closable](target T, err error) T {
	closables = append(closables, target)

	return load(target, err)
}

func load[T any](target T, err error) T {
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
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	for _, runnable := range runnables {
		wg.Add(1)

		go func(runnable Runnable) {
			if err := runnable.Run(ctx, &wg); err != nil {
				logger.Errorf("[MAIN] %s run failed: %s", reflect.TypeOf(runnable), err)
				cancel()
			}
		}(runnable)
	}

	waitTermination(ctx)
	cancel()
	wg.Wait()
}

func waitTermination(ctx context.Context) {
	defer exit()

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		logger.Infoln("[MAIN] received context done")
	case s := <-terminationChan:
		logger.Infof("[MAIN] received signal: %s", s)
	}
}
