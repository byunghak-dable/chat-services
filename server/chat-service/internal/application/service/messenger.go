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
	lockByRoom     *sync.Map
	roomById       map[string]map[driver.MessengerClient]struct{}
}

func NewMessenger(logger driven.Logger, broker driven.MessageBroker, messageStore driver.Message) *Messenger {
	messenger := &Messenger{
		logger,
		broker,
		messageStore,
		new(sync.Map),
		make(map[string]map[driver.MessengerClient]struct{}),
	}
	broker.Subscribe(messenger)

	return messenger
}

func (m *Messenger) OnReceive(message *dto.Message) {
	roomId := message.RoomId

	m.withRLock(roomId, func() {
		room := m.roomById[roomId]

		for participant := range room {
			go func(participant driver.MessengerClient) {
				if err := participant.Send(message); err != nil {
					m.logger.Errorf("[MESSENGER] failed to send message: %s", err)
				}
			}(participant)
		}
	})
}

func (m *Messenger) Join(client driver.MessengerClient) {
	roomId := client.RoomId()

	m.withLock(roomId, func() {
		if _, ok := m.roomById[roomId]; !ok {
			m.roomById[roomId] = make(map[driver.MessengerClient]struct{})
		}
		m.roomById[roomId][client] = struct{}{}
	})
}

func (m *Messenger) Leave(client driver.MessengerClient) {
	roomId := client.RoomId()

	m.withLock(roomId, func() {
		room := m.roomById[roomId]

		delete(room, client)

		if len(room) == 0 {
			delete(m.roomById, roomId)
			m.lockByRoom.Delete(roomId)
		}
	})
}

func (m *Messenger) Send(message *dto.Message) error {
	if err := m.messageService.Save(message); err != nil {
		return err
	}
	return m.broker.Publish(message)
}

func (m *Messenger) withLock(roomId string, action func()) {
	lock := m.getLock(roomId)
	lock.Lock()
	defer lock.Unlock()

	action()
}

func (m *Messenger) withRLock(roomId string, action func()) {
	lock := m.getLock(roomId)
	lock.RLock()
	defer lock.RUnlock()

	action()
}

func (m *Messenger) getLock(roomId string) *sync.RWMutex {
	mutex, _ := m.lockByRoom.LoadOrStore(roomId, &sync.RWMutex{})

	return mutex.(*sync.RWMutex)
}
