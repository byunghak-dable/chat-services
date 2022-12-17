package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gorilla/websocket"
)

type messageObj struct {
	RoomIdx uint   `json:"room_idx"`
	UserIdx uint   `json:"user_idx"`
	Message string `json:"message"`
}

func main() {
	wg := &sync.WaitGroup{}

	arr := [2]*Client{}
	for i := range arr {
		c := NewClient()
		go c.read(wg)
		arr[i] = c
	}

	for {
		var msg string
		fmt.Scanln(&msg)
		if msg == "exit" {
			break
		}
		for _, c := range arr {
			c.conn.WriteJSON(&messageObj{
				RoomIdx: 1,
				UserIdx: 1,
				Message: msg,
			})
		}
	}

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan

	for _, c := range arr {
		c.conn.Close()
	}
	wg.Wait()
	log.Println("waiting finished")
}

type Client struct {
	conn   *websocket.Conn
	doneCh chan struct{}
}

func NewClient() *Client {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8082/api/v1/chat?user_idx=1&room_idx=1", nil)
	if err != nil {
		log.Println("dial:", err)
	}
	return &Client{
		conn: conn,
	}
}

func (c *Client) write(message string) {
	err := c.conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write:", err)
		return
	}
}

func (c *Client) read(wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("recv: %s", msg)
	}
}
