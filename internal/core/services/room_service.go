package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cristianhuang/omuricechatserver/internal/core/domain"
	"github.com/cristianhuang/omuricechatserver/internal/core/ports/input"
)

type RoomService struct{}

func NewRoomService() *RoomService {
	return &RoomService{}
}

func (r *RoomService) HandleConnection(conn input.Connection, room *domain.Room) {

	client := &domain.Client{Conn: conn, Send: make(chan []byte, 256)}
	room.Register <- client

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
		if msg.ID == "" || msg.Message == nil {
			errMsg, _ := json.Marshal(map[string]string{
				"error": "missing required fields: id, message",
			})
			conn.WriteMessage(1, errMsg)
			continue
		}

		room.Broadcast <- domain.MessageSend{ID: msg.ID, Message: msg.Message, SentAt: time.Now().String()}
	}
	room.Unregister <- client
}

type RoomServiceInterface interface {
	HandleConnection(conn input.Connection, room *domain.Room)
}
