package service

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/entity"
	"chat-service/internal/port/driver"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type MockClient struct {
	sendCount *uint32
	roomId    string
	userId    string
}

func (m *MockClient) RoomId() string {
	return m.roomId
}

func (m *MockClient) UserId() string {
	return m.userId
}

func (m *MockClient) Send(message dto.Message) error {
	atomic.AddUint32(m.sendCount, 1)
	return nil
}

func TestThreadSafety(t *testing.T) {
	// given
	var clients []*MockClient
	var sendCount uint32
	roomById := make(map[string]*entity.LiveRoom)
	roomManager := &RoomManager{rooms: roomById}

	for i := range 100 {
		for j := 0; j < 100; j++ {
			clients = append(clients, &MockClient{
				sendCount: &sendCount,
				roomId:    fmt.Sprintf("room-%d", i),
				userId:    fmt.Sprintf("user-%d", j),
			})
		}
	}

	// when
	var wg sync.WaitGroup

	for _, client := range clients {
		wg.Add(1)

		go func(c driver.MessengerClient) {
			defer wg.Done()

			roomManager.Join(c)
			_ = roomManager.Broadcast(dto.Message{RoomId: c.RoomId()})
			time.Sleep(10 * time.Millisecond)
			roomManager.Leave(c)
		}(client)
	}

	wg.Wait()

	// then
	expectedRooms := 0
	actualRooms := len(roomById)

	println(sendCount)
	if actualRooms != expectedRooms {
		t.Errorf("Expected %d rooms, but got %d", expectedRooms, actualRooms)
	}
}
