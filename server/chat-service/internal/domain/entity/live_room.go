package entity

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/port/driver"
	"sync"
)

type LiveRoom struct {
	participants map[driver.MessengerClient]struct{}
	id           string
	mu           sync.RWMutex
}

func NewLiveRoom(id string) *LiveRoom {
	return &LiveRoom{
		participants: make(map[driver.MessengerClient]struct{}),
		id:           id,
	}
}

func (r *LiveRoom) Join(client driver.MessengerClient) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.participants[client] = struct{}{}
}

func (r *LiveRoom) Leave(client driver.MessengerClient) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.participants, client)
}

func (r *LiveRoom) Broadcast(message dto.Message) error {
	errChan := r.broadcast(message)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *LiveRoom) IsEmpty() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.participants) == 0
}

func (r *LiveRoom) Id() string {
	return r.id
}

func (r *LiveRoom) broadcast(message dto.Message) <-chan error {
	var wg sync.WaitGroup
	errChan := make(chan error)

	r.mu.RLock()
	defer r.mu.RUnlock()

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
