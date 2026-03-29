package services

import (
	"bytes"
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
		decoder := json.NewDecoder(bytes.NewReader(m))
		decoder.DisallowUnknownFields()
		if err = decoder.Decode(&msg); err != nil {
			errMsg, _ := json.Marshal(map[string]string{
				"error": "invalid format message: " + err.Error(),
			})
			conn.WriteMessage(1, errMsg)
			continue
		}
		if msg.SenderID == "" || msg.Message == nil {
			errMsg, _ := json.Marshal(map[string]string{
				"error": "missing required fields: sender_id, message",
			})
			conn.WriteMessage(1, errMsg)
			continue
		}

		h.hub.Broadcast <- domain.MessageSend{SenderID: msg.SenderID, Message: msg.Message, SentAt: time.Now().String()}
	}
	h.hub.Unregister <- client
}
