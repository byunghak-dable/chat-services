package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"messenger-service/internal/adapter/driven"
	"messenger-service/internal/adapter/driving"
	"messenger-service/internal/adapter/driving/grpc"
	"messenger-service/internal/adapter/driving/rest"
	"messenger-service/internal/application/service"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"syscall"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
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
	kafkaProducer, producerErr := driven.NewKafkaProducer(&kafka.ConfigMap{
		"bootstrap.servers": getKafkaServers(),
		"client.id":         "TEST_CLIENT_ID",
		"acks":              "all",
	})

	defer quit(kafkaProducer)

	if producerErr != nil {
		logger.Error(producerErr)
		return
	}

	messengerService := service.NewMessengerService(logger, kafkaProducer, service.NewRoomService())

	restApp := rest.New(logger, messengerService)
	grpcApp := grpc.New(logger, messengerService)
	kafkaConsumer, consumerErr := driving.NewKafkaConsumer(&kafka.ConfigMap{
		"bootstrap.servers": getKafkaServers(),
		"auto.offset.reset": "smallest",
	}, messengerService)

	defer quit(restApp, grpcApp, kafkaProducer)

	if consumerErr != nil {
		logger.Error(consumerErr)
		return
	}

	go restApp.Run(os.Getenv("REST_PORT"))
	go grpcApp.Run(os.Getenv("GRPC_PORT"))
	kafkaConsumer.Run()

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}

func getKafkaServers() string {
	servers := []string{
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_1_HOST"), os.Getenv("KAFKA_1_PORT")),
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_2_HOST"), os.Getenv("KAFKA_2_PORT")),
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_3_HOST"), os.Getenv("KAFKA_3_PORT")),
	}

	return strings.Join(servers, ",")
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
