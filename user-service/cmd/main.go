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

var logger = log.New()
var (
	mysqlDb *mysql.Mysql
	redisDb *redis.Redis
)
var (
	restServer *rest.Rest
	grpcServer *grpc.Grpc
)

// env
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
}

// DB
func init() {
	var mysqlErr, redisErr error
	mysqlDb, mysqlErr = mysql.New(logger, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	redisDb, redisErr = redis.New(logger, net.JoinHostPort(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASSWORD"), 0)

	if mysqlErr != nil || redisErr != nil {
		shutdown(redisDb, mysqlDb)
		logger.Fatalf("database connection failure\nmysql: %s\nredis: %s", mysqlErr, redisErr)
	}
}

// servers
func init() {
	userRepo := repository.NewUserRepo(logger, mysqlDb)
	userApp := application.NewUserApp(logger, userRepo)
	restServer = rest.New(logger, userApp)
}

func main() {
	defer gracefulShutdown()
	go restServer.Run(os.Getenv("REST_PORT"))
}

func gracefulShutdown() {
	defer shutdown(restServer, mysqlDb, redisDb)

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}

func shutdown(targets ...interface{ Close() error }) {
	for _, target := range targets {
		if !reflect.ValueOf(target).IsNil() {
			if err := target.Close(); err != nil {
				logger.Errorf("%s closing failed: %s", reflect.TypeOf(target), err)
			}
		}
	}
}
