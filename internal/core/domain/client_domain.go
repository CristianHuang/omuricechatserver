package domain

import (
	"time"

	"github.com/cristianhuang/omuricechatserver/internal/core/ports/input"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn input.Connection
	Send chan []byte
}

func (c *Client) Writer() {
	for message := range c.Send {
		c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
			return
		}
	}
}
