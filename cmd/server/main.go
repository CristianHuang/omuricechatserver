package main

import (
	"github.com/cristianhuang/omuricechatserver/internal/adapters/handlers/websocket"
	"github.com/cristianhuang/omuricechatserver/internal/core/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	roomService := services.NewRoomService()
	wsHandler := websocket.NewWsHandler(roomService)

	r.GET("/chat/room/:id", wsHandler.HandleRoom)

	r.Run(":8080")
}
