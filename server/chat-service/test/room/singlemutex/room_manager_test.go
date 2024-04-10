package singlemutex

// Performance comparison between "Room Manager of Production" and "Room Manager using single Mutex"

import (
	"chat-service/internal/application/abstraction"
	"chat-service/internal/application/dto"
	"chat-service/internal/domain/service"
	"chat-service/internal/port/driver"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type MockClient struct {
	count  *int32
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
	atomic.AddInt32(m.count, 1)
	time.Sleep(time.Millisecond)
	return nil
}

const (
	benchRoomCount = 1000
	/**
	 * Because the number of calls may vary due to differences in "client.Send" locking methods.
	 * the number of chat participants in each room is limited to one.
	 */
	benchParticipantCount = 1
	benchBroadcastCount   = 10
)

func BenchmarkLiveRoomManager(b *testing.B) {
	stressTest(service.NewRoomManager())
}

func BenchmarkSingleMutexRoomManager(b *testing.B) {
	stressTest(NewRoomManager())
}

func stressTest(roomManager abstraction.RoomManager) {
	// given
	var clients []*MockClient
	var count int32

	for i := range benchRoomCount {
		for j := range benchParticipantCount {
			clients = append(clients, &MockClient{
				roomId: fmt.Sprintf("room-bench-%d", i),
				userId: fmt.Sprintf("user-bench-%d", j),
				count:  &count,
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
			defer roomManager.Leave(c)

			for range benchBroadcastCount {
				_ = roomManager.Broadcast(dto.Message{RoomId: c.RoomId()})
			}
		}(client)
	}

	wg.Wait()
}
