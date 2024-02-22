package main

import (
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/messenger-service/internal/adapter/driven"
	"github.com/widcraft/messenger-service/internal/adapter/driving/grpc"
	"github.com/widcraft/messenger-service/internal/adapter/driving/rest"
	"github.com/widcraft/messenger-service/internal/application/service"
)

var logger = log.New()

type Exitable interface {
	OnExit() error
}

// env
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	kafkaProducer, kafkaErr := driven.NewKafkaProducer(makeKafkaProducerConfig())

	defer exit(kafkaProducer)

	if kafkaErr != nil {
		logger.Error(kafkaErr)
		return
	}

	messengerService := service.NewMessengerService(logger, kafkaProducer)

	restApp := rest.New(logger, messengerService)
	grpcApp := grpc.New(logger, messengerService)

	defer exit(restApp, grpcApp)

	go restApp.Run(os.Getenv("REST_PORT"))
	go grpcApp.Run(os.Getenv("GRPC_PORT"))

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}

func makeKafkaProducerConfig() *driven.KafkaProducerConfig {
	return &driven.KafkaProducerConfig{
		Topic:    os.Getenv("KAFKA_CHAT_TOPIC"),
		ClientId: "testing",
		Nodes: []driven.KafkaNode{
			{Host: os.Getenv("KAFKA_1_HOST"), Port: os.Getenv("KAFKA_1_PORT")},
			{Host: os.Getenv("KAFKA_2_HOST"), Port: os.Getenv("KAFKA_2_PORT")},
			{Host: os.Getenv("KAFKA_3_HOST"), Port: os.Getenv("KAFKA_3_PORT")},
		},
	}
}

func exit(exitables ...Exitable) {
	for _, exitable := range exitables {
		if reflect.ValueOf(exitable).IsNil() {
			continue
		}

		err := exitable.OnExit()
		if err != nil {
			logger.Errorf("%s exiting failed: %s", reflect.TypeOf(exitable), err)
		}
	}
}
