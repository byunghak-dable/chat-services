package service

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/entity"
	"chat-service/internal/port/driver"
	"fmt"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var count uint64 = 0

type MockClient struct {
	roomId string
	userId string
}

func (m *MockClient) RoomId() string {
	return m.roomId
}

func (m *MockClient) UserId() string {
	return m.userId
}

func (m *MockClient) Send(message dto.Message) error {
	atomic.AddUint64(&count, 1)
	return nil
}

func TestThreadSafety(t *testing.T) {
	// given
	roomById := make(map[string]*entity.LiveRoom)
	roomManager := &RoomManager{roomById: roomById}
	var clients []*MockClient

	for i := range 500 {
		for j := 0; j < 500; j++ {
			clients = append(clients, &MockClient{fmt.Sprintf("room-%d", i), fmt.Sprintf("user-%d", j)})
		}
	}

	// when
	var wg sync.WaitGroup

	for _, client := range clients {
		wg.Add(1)
		go func(c driver.MessengerClient) {
			defer wg.Done()

			time.Sleep(rand.N(50 * time.Millisecond))
			roomManager.Join(c)
			roomManager.Leave(c)
		}(client)
	}

	wg.Wait()

	// then
	expectedRooms := 0
	actualRooms := len(roomById)

	println(count)
	if actualRooms != expectedRooms {
		t.Errorf("Expected %d rooms, but got %d", expectedRooms, actualRooms)
	}
}
