package domain

import (
	"encoding/json"
)

type Hub struct {
	Clients    map[*Client]bool
	Id         chan int
	Broadcast  chan MessageSend
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan MessageSend),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				client.Conn.Close()
			}
		case message := <-h.Broadcast:
			msg, _ := json.Marshal(message)
			for client := range h.Clients {
				select {
				case client.Send <- msg:
				default:
					close(client.Send)
					client.Conn.Close()
					delete(h.Clients, client)
				}
			}
		}
	}
}
