package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/widcraft/chat-service/internal/adapter/grpc/chat/pb"
	"google.golang.org/grpc"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go runClient(wg)
	}
	wg.Wait()
}

func runClient(wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to: %s", err)
	}
	defer conn.Close()
	client := pb.NewChatClient(conn)
	stream, err := client.Connect(context.Background())
	if err != nil {
		log.Fatalf("response error: %s", err)
	}

	go receiveMessage(stream)
	sendMessage(stream)
}

func sendMessage(stream pb.Chat_ConnectClient) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err := stream.Send(&pb.ChatReq{
			Type: &pb.ChatReq_Join{
				Join: &pb.JoinReq{
					RoomIdx: 1,
					UserIdx: 1,
				},
			},
		})
		if err != nil {
			log.Fatalf("send to server err: %s", err)
			break
		}
	}
}

func receiveMessage(stream pb.Chat_ConnectClient) {
	for {
		message, err := stream.Recv()
		if err != nil {
			log.Fatalf("client stream error: %s", err)
			break
		}
		fmt.Println(message)
	}
}
