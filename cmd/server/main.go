package main

import (
	"github.com/cristianhuang/omuricechatserver/internal/adapters/handlers/websocket"
	"github.com/cristianhuang/omuricechatserver/internal/core/domain"
	"github.com/cristianhuang/omuricechatserver/internal/core/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	hub := domain.NewHub()
	go hub.Run()

	hubService := services.NewHubService(hub)
	wsHandler := websocket.NewWsHandler(hubService)

	r.GET("/ws", wsHandler.Ws)

	r.Run(":8080")
}
