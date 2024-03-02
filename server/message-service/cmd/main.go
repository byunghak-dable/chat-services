package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"message-service/internal/application/service"
)

// env
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Print("testing")
	service.New{}
}
