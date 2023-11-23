package main

import (
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/user-service/internal/adapter/grpc"
	"github.com/widcraft/user-service/internal/adapter/repository"
	"github.com/widcraft/user-service/internal/adapter/rest"
	"github.com/widcraft/user-service/internal/application"
	"github.com/widcraft/user-service/pkg/db"
)

var logger = log.New()

type Closable interface {
	Close() error
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	mysqlConfig := db.MysqlConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}
	redisConfig := db.RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Db:       0,
	}
	mysqlDb, mysqlErr := db.NewMysql(mysqlConfig)
	redisDb, redisErr := db.NewRedis(redisConfig)

	defer shutdown(mysqlDb, redisDb)

	if mysqlErr != nil || redisErr != nil {
		logger.Error(mysqlErr, redisErr)
		return
	}

	userRepo := repository.NewUserRepo(logger, mysqlDb)
	userApp := application.NewUserApp(logger, userRepo)
	restServer := rest.New(logger, userApp)
	grpcServer := grpc.New(logger, userApp)

	defer shutdown(restServer, grpcServer)

	go restServer.Run(os.Getenv("REST_PORT"))
	go grpcServer.Run(os.Getenv("GRPC_PORT"))

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}

func shutdown(closables ...Closable) {
	for _, closable := range closables {
		if reflect.ValueOf(closable).IsNil() {
			continue
		}
		if err := closable.Close(); err != nil {
			logger.Errorf("%s closing failed: %s", reflect.TypeOf(closable), err)
		}
	}
}
