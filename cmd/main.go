package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/weed082/chat-server/internal/adapter/driven/repository"
	"github.com/weed082/chat-server/internal/adapter/driven/repository/mysql"
	"github.com/weed082/chat-server/internal/adapter/driven/repository/redis"
	"github.com/weed082/chat-server/internal/adapter/driver/rest"
	"github.com/weed082/chat-server/internal/application"
)

var logger = log.New(os.Stdout, "LOG", log.LstdFlags|log.Llongfile)

var (
	restServer *rest.Rest
	mysqlDb    *mysql.Mysql
	redisDb    *redis.Redis
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
	// init mysql
	sqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	mysqlDb = mysql.New(logger, sqlDsn)
	// init redis
	redisHost := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	redisDb = redis.New(logger, redisHost, os.Getenv("REDIS_PASSWORD"), 0)
}

func main() {
	defer gracefulShutdown()
	runRest()
}

func runRest() {
	userRepo := repository.NewUserRepo(logger, mysqlDb)
	userApp := application.NewUserApp(logger, userRepo)
	restServer = rest.New(logger, userApp)
	go restServer.Run(os.Getenv("REST_PORT"))
}

func gracefulShutdown() {
	defer mysqlDb.Close()
	defer restServer.Stop()

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}
