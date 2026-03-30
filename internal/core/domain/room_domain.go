package domain

import "encoding/json"

type Room struct {
	ID         string
	Clients    map[*Client]bool
	Broadcast  chan MessageSend
	Register   chan *Client
	Unregister chan *Client
}

func NewRoom(id string) *Room {
	return &Room{
		ID:         id,
		Broadcast:  make(chan MessageSend),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Room) Run() {
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
					delete(h.Clients, client)
					close(client.Send)
					client.Conn.Close()
				}
			}
		}
	}
}
