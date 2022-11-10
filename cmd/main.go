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
	"github.com/weed082/chat-server/internal/adapter/driver/rest"
	"github.com/weed082/chat-server/internal/application"
)

var (
	logger = log.New(os.Stdout, "LOG", log.LstdFlags|log.Llongfile)
)

var (
	restServer *rest.Rest
	db         *mysql.Mysql
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file")
	}
	sqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	db = mysql.New(logger, sqlDsn)
}

func main() {
	runRest()
	gracefulShutdown()
}

func runRest() {
	userRepo := repository.NewUserRepo(logger, db)
	userApp := application.NewUserApp(logger, userRepo)
	restServer = rest.New(logger, userApp)
	go restServer.Run("3000")
}

func gracefulShutdown() {
	defer db.Close()
	defer restServer.Stop()

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}
