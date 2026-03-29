package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cristianhuang/omuricechatserver/internal/core/domain"
	"github.com/cristianhuang/omuricechatserver/internal/core/ports/input"
)

type HubService struct {
	hub *domain.Hub
}

func NewHubService(hub *domain.Hub) *HubService {
	return &HubService{hub: hub}
}

func (h *HubService) HandleConnection(conn input.Connection) {
	client := &domain.Client{Conn: conn, Send: make(chan []byte, 256)}
	h.hub.Register <- client

	go client.Writer()

	for {
		var msg domain.Message
		m, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message: ", err)
			break
		}
		err = json.Unmarshal(m, &msg)
		if err != nil {
			fmt.Println("Error unmarshing message: ", err)
			break
		}

		h.hub.Broadcast <- domain.MessageSend{SenderID: msg.SenderID, Message: msg.Message, SentAt: time.Now().String()}
	}
	h.hub.Unregister <- client
}
