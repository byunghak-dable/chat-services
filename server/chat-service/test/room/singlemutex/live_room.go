package singlemutex

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/port/driver"
	"sync"
)

type LiveRoom struct {
	participants map[driver.MessengerClient]struct{}
	id           string
}

func NewLiveRoom(id string) *LiveRoom {
	return &LiveRoom{
		participants: make(map[driver.MessengerClient]struct{}),
		id:           id,
	}
}

func (r *LiveRoom) Join(client driver.MessengerClient) {
	r.participants[client] = struct{}{}
}

func (r *LiveRoom) Leave(client driver.MessengerClient) {
	delete(r.participants, client)
}

func (r *LiveRoom) Broadcast(message dto.Message) error {
	for err := range r.broadcast(message) {
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *LiveRoom) IsEmpty() bool {
	return len(r.participants) == 0
}

func (r *LiveRoom) broadcast(message dto.Message) <-chan error {
	var wg sync.WaitGroup
	errChan := make(chan error)

	for participant := range r.participants {
		wg.Add(1)

		go func(participant driver.MessengerClient) {
			defer wg.Done()
			errChan <- participant.Send(message)
		}(participant)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	return errChan
}
