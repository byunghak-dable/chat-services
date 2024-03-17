package service

import (
	"chat-service/internal/application/dto"
	"chat-service/internal/port/driven"
	"chat-service/internal/port/driver"
	"sync"
)

type Messenger struct {
	logger         driven.Logger
	broker         driven.MessageBroker
	messageService driver.Message
	lockByRoom     sync.Map
	roomById       map[string][]driver.MessengerClient
}

func NewMessenger(logger driven.Logger, broker driven.MessageBroker, messageStore driver.Message) *Messenger {
	messenger := &Messenger{logger, broker, messageStore, sync.Map{}, make(map[string][]driver.MessengerClient)}
	broker.Subscribe(messenger)

	return messenger
}

func (m *Messenger) OnReceive(message *dto.Message) {
	roomIdx := message.RoomId

	m.withRLock(roomIdx, func() {
		room := m.roomById[roomIdx]

		for _, participant := range room {
			go func(participant driver.MessengerClient) {
				if err := participant.Send(message); err != nil {
					m.logger.Errorf("[MESSENGER] failed to send message: %s", err)
				}
			}(participant)
		}
	})
}

func (m *Messenger) Join(client driver.MessengerClient) {
	roomIdx := client.RoomId()

	m.withLock(roomIdx, func() {
		m.roomById[roomIdx] = append(m.roomById[roomIdx], client)
	})

}

func (m *Messenger) Leave(client driver.MessengerClient) {
	roomIdx := client.RoomId()

	m.withLock(roomIdx, func() {
		participants := m.roomById[roomIdx]

		for i, participant := range m.roomById[roomIdx] {
			if client != participant {
				continue
			}

			m.roomById[roomIdx] = append(participants[:i], participants[i+1:]...)

			if len(m.roomById[roomIdx]) == 0 {
				delete(m.roomById, roomIdx)
				m.lockByRoom.Delete(roomIdx)
			}
		}
	})
}

func (m *Messenger) Send(message *dto.Message) error {
	if err := m.messageService.Save(message); err != nil {
		return err
	}
	return m.broker.Publish(message)
}

func (m *Messenger) withLock(roomIdx string, action func()) {
	lock := m.getLock(roomIdx)
	lock.Lock()
	defer lock.Unlock()

	action()
}

func (m *Messenger) withRLock(roomIdx string, action func()) {
	lock := m.getLock(roomIdx)
	lock.RLock()
	defer lock.RUnlock()

	action()
}

func (m *Messenger) getLock(roomIdx string) *sync.RWMutex {
	mutex, _ := m.lockByRoom.LoadOrStore(roomIdx, &sync.RWMutex{})

	return mutex.(*sync.RWMutex)
}
