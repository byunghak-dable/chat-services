package main

import (
	"net"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/widcraft/user-service/internal/adapter/grpc"
	"github.com/widcraft/user-service/internal/adapter/repository"
	"github.com/widcraft/user-service/internal/adapter/repository/mysql"
	"github.com/widcraft/user-service/internal/adapter/repository/redis"
	"github.com/widcraft/user-service/internal/adapter/rest"
	"github.com/widcraft/user-service/internal/application"
)

type Closable interface {
	Close() error
}

// env
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
}

func main() {
	logger := log.New()
	mysqlDb, mysqlErr := mysql.New(logger, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	redisDb, redisErr := redis.New(logger, net.JoinHostPort(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASSWORD"), 0)
	defer shutdown(logger, mysqlDb, redisDb)

	if mysqlErr != nil || redisErr != nil {
		logger.Error(mysqlErr, redisErr)
		return
	}

	userRepo := repository.NewUserRepo(logger, mysqlDb)
	userApp := application.NewUserApp(logger, userRepo)
	restServer := rest.New(logger, userApp)
	grpcServer := grpc.New(logger, userApp)
	defer shutdown(logger, restServer, grpcServer)

	go restServer.Run(os.Getenv("REST_PORT"))
	go grpcServer.Run(os.Getenv("GRPC_PORT"))

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}

func shutdown(logger log.FieldLogger, closables ...Closable) {
	for _, target := range closables {
		if reflect.ValueOf(target).IsNil() {
			continue
		}
		if err := target.Close(); err != nil {
			logger.Errorf("%s closing failed: %s", reflect.TypeOf(target), err)
		}
	}
}
