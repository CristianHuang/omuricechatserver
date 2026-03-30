package domain

import (
	"sync"
)

type Hub struct {
	Rooms map[string]*Room
	mu    sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*Room),
	}
}

func (h *Hub) CreateRoom(id string) *Room {
	h.mu.Lock()
	defer h.mu.Unlock()

	room := NewRoom(id)

	h.Rooms[id] = room
	go room.Run()
	return room
}

func (h *Hub) GetRoom(id string) (*Room, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()

	room, ok := h.Rooms[id]
	return room, ok
}
