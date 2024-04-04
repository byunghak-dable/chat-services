package entity

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/port/driver"
	"fmt"
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
	r.mu.RLock()
	defer r.mu.RUnlock()

	count := len(r.participants)
	errChan := make(chan error, count)

	defer close(errChan)

	for participant := range r.participants {
		go func(participant driver.MessengerClient) {
			errChan <- participant.Send(message)
		}(participant)
	}

	var errs []error

	for range count {
		if err := <-errChan; err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%v", errs)
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
