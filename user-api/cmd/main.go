package main

import (
	"net"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/weed082/user-api/internal/adapter/driven/repository"
	"github.com/weed082/user-api/internal/adapter/driven/repository/mysql"
	"github.com/weed082/user-api/internal/adapter/driven/repository/redis"
	"github.com/weed082/user-api/internal/adapter/driver/grpc"
	"github.com/weed082/user-api/internal/adapter/driver/rest"
	"github.com/weed082/user-api/internal/application"
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

// init env
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
}

// init database
func init() {
	var mysqlErr, redisErr error
	mysqlDb, mysqlErr = mysql.New(logger, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	redisDb, redisErr = redis.New(logger, net.JoinHostPort(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASSWORD"), 0)

	if mysqlErr != nil || redisErr != nil {
		shutdown(redisDb, mysqlDb)
		logger.Fatalf("database connection failure\nmysql: %s\nredis: %s", mysqlErr, redisErr)
	}
}

// init servers
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

func shutdown(args ...interface{ Close() }) {
	for _, arg := range args {
		if !reflect.ValueOf(arg).IsNil() {
			arg.Close()
		}
	}
}
