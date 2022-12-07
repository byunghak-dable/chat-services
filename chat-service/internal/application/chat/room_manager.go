package chat

import (
	"errors"
	"strconv"
	"strings"

	"github.com/widcraft/chat-service/internal/domain/dto"
	"github.com/widcraft/chat-service/internal/port"
)

// TOOD: add mutex
type roomManager struct {
	rooms map[uint][]port.ChatClient
}

func NewRoomManager() *roomManager {
	return &roomManager{
		rooms: make(map[uint][]port.ChatClient),
	}
}

func (manager *roomManager) add(roomIdx uint, client port.ChatClient) {
	room, ok := manager.rooms[roomIdx]
	if ok {
		manager.rooms[roomIdx] = append(room, client)
		return
	}
	manager.rooms[roomIdx] = []port.ChatClient{client}
}

func (manager *roomManager) quit(roomIdx uint, client port.ChatClient) error {
	room, ok := manager.rooms[roomIdx]
	if !ok {
		return errors.New("no existing chat room roomIdx")
	}
	for i, roomClient := range room {
		if client == roomClient {
			manager.rooms[roomIdx] = append(room[:i], room[i+1:]...)
			return nil
		}
	}
	return errors.New("no client in chat room")
}

func (manager *roomManager) sendMessage(message dto.MessageDto) error {
	room, ok := manager.rooms[message.RoomIdx]
	if !ok {
		return errors.New("no existing chat room roomIdx")
	}

	failedClients := []string{}
	for _, client := range room {
		if err := client.SendMessage(message); err != nil {
			failedClients = append(failedClients, strconv.FormatUint(uint64(client.GetUserIdx()), 10))
		}
	}

	if len(failedClients) > 0 {
		return errors.New("some clients failed to send message: " + strings.Join(failedClients, ", "))
	}
	return nil
}
