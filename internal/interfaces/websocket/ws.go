package websocket

import (
	"fmt"
	"net/http"

	"github.com/cristianhuang/omuricechatserver/internal/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Ws(hub *domain.Hub, c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrade:", err)
		return
	}
	client := &domain.Client{Conn: conn, Send: make(chan []byte, 256)}
	hub.Register <- client

	go client.Writer()

	fmt.Println("Cliente conectado!")

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		hub.Broadcast <- message
	}
	hub.Unregister <- client
}
