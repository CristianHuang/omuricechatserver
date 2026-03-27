package main

import (
	"github.com/cristianhuang/omuricechatserver/internal/core/domain"
	"github.com/cristianhuang/omuricechatserver/internal/interfaces/websocket"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	hub := domain.NewHub()
	go hub.Run()

	r.GET("/ws", func(c *gin.Context) { websocket.Ws(hub, c) })

	r.Run(":8080")
}
