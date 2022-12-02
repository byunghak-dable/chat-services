package main

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/widcraft/chat-service/external/workerpool"
)

var logger = log.New()
var (
	wg       = &sync.WaitGroup{}
	chatPool = workerpool.New(wg, 1)
)

func main() {
	fmt.Println("chat service main")
}

func gracefulShutdown() {
	// waiting
	defer wg.Wait()
	// closing
	defer shutdown(chatPool)

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
}

func shutdown(targets ...interface{ Close() }) {
	for _, target := range targets {
		if !reflect.ValueOf(target).IsNil() {
			target.Close()
		}
	}
}
