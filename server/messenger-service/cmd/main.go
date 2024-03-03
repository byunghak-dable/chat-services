package main

import (
	"context"
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

// env
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	kafkaProducer, producerErr := driven.NewKafkaProducer(getKafkaProducerConfig())

	defer quit(kafkaProducer)

	if producerErr != nil {
		logger.Error(producerErr)
		return
	}

	messengerService := service.NewMessengerService(logger, kafkaProducer, service.NewRoomService())

	kafkaConsumer, consumerErr := driving.NewKafkaConsumer(logger, getKafkaConsumerConfig(), messengerService, []string{"topic1"})
	restApp := rest.New(logger, messengerService, os.Getenv("REST_PORT"))
	grpcApp := grpc.New(logger, messengerService, os.Getenv("GRPC_PORT"))

	defer quit(restApp, grpcApp, kafkaProducer)

	if consumerErr != nil {
		logger.Error(consumerErr)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	run(cancel, restApp, grpcApp, kafkaConsumer)
	handleTermination(ctx)
}

func getKafkaProducerConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": getKafkaServers(),
		"client.id":         "TEST_CLIENT_ID",
		"acks":              "all",
	}
}

func getKafkaConsumerConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": getKafkaServers(),
		"group.id":          "TEST_GROUP_ID",
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

func run(cancel context.CancelFunc, runnables ...interface{ Run() error }) {
	for _, runnable := range runnables {
		go func() {
			err := runnable.Run()

			if err != nil {
				logger.Errorf("%s failed: %s", reflect.TypeOf(runnable), err)
				cancel()
			}
		}()
	}

}

func quit(closables ...interface{ Close() error }) {
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
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	select {
	case <-ctx.Done():
		fmt.Println("Received context done")
	case s := <-terminationChan:
		fmt.Println("Received signal:", s)
	}
}
