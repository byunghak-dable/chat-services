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
	lock         sync.RWMutex
}

func NewLiveRoom(id string) *LiveRoom {
	return &LiveRoom{
		participants: make(map[driver.MessengerClient]struct{}),
		id:           id,
	}
}

func (r *LiveRoom) Join(client driver.MessengerClient) {
	r.withLock(func() {
		r.participants[client] = struct{}{}
	})
}

func (r *LiveRoom) Leave(client driver.MessengerClient) {
	r.withLock(func() {
		delete(r.participants, client)
	})
}

func (r *LiveRoom) Broadcast(message dto.Message) error {
	var count int
	errChan := make(chan error)

	defer close(errChan)

	r.withRLock(func() {
		count = len(r.participants)

		for participant := range r.participants {
			go func(participant driver.MessengerClient) {
				errChan <- participant.Send(message)
			}(participant)
		}
	})

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
	var isEmpty bool

	r.withRLock(func() {
		isEmpty = len(r.participants) == 0
	})

	return isEmpty
}

func (r *LiveRoom) ID() string {
	return r.id
}

func (r *LiveRoom) withLock(action func()) {
	r.lock.Lock()
	defer r.lock.Unlock()

	action()
}

func (r *LiveRoom) withRLock(action func()) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	action()
}
