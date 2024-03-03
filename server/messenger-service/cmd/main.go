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
	"strings"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Runnable interface {
	Run() error
}

type Closable interface {
	Close() error
}

var logger = log.New()

// env
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	topic := os.Getenv("KAFKA_CHAT_TOPIC")
	kafkaProducer, producerErr := driven.NewKafkaProducer(getKafkaProducerConfig(), topic)

	defer quit(kafkaProducer)

	if producerErr != nil {
		logger.Error(producerErr)
		return
	}

	messengerService := service.NewMessengerService(logger, kafkaProducer, service.NewRoomService())

	kafkaConsumer, consumerErr := driving.NewKafkaConsumer(logger, getKafkaConsumerConfig(), messengerService, []string{topic})
	restApp := rest.New(logger, messengerService, os.Getenv("REST_PORT"))
	grpcApp := grpc.New(logger, messengerService, os.Getenv("GRPC_PORT"))

	defer quit(restApp, grpcApp, kafkaProducer)

	if consumerErr != nil {
		logger.Error(consumerErr)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	run(cancel, kafkaConsumer, restApp, grpcApp)
	handleTermination(ctx)
}

func getKafkaProducerConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": getKafkaServers(),
		"client.id":         os.Getenv("KAFKA_CLIENT_ID"),
		"acks":              "all",
	}
}

func getKafkaConsumerConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": getKafkaServers(),
		"group.id":          os.Getenv("KAFKA_GROUP_ID"), // TODO: need to add suffix for scale out
		"auto.offset.reset": "smallest",
	}
}

func getKafkaServers() string {
	servers := []string{
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_1_HOST"), os.Getenv("KAFKA_1_PORT")),
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_2_HOST"), os.Getenv("KAFKA_2_PORT")),
		fmt.Sprintf("%s:%s", os.Getenv("KAFKA_3_HOST"), os.Getenv("KAFKA_3_PORT")),
	}

	return strings.Join(servers, ",")
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
