package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/weed082/chat-server/internal/adapter/driving/rest"
)

var (
	logger = log.New(os.Stdout, "LOG", log.LstdFlags|log.Llongfile)
)

var (
	restServer *rest.Rest
)

func main() {
	runRest()
	gracefulShutdown()
}

func runRest() {
	restServer = rest.New(logger)
	go restServer.Run("3000")
}

func gracefulShutdown() {
	defer restServer.Stop()

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}
