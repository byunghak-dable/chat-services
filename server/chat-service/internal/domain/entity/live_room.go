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
	var errs []error
	var lock sync.Mutex

	r.withRLock(func() {
		for participant := range r.participants {
			go func(participant driver.MessengerClient) {
				if err := participant.Send(message); err != nil {
					lock.Lock()
					errs = append(errs, err)
					lock.Unlock()
				}
			}(participant)
		}
	})

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
