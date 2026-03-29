package websocket

import (
	"fmt"
	"net/http"

	"github.com/cristianhuang/omuricechatserver/internal/core/ports/input"
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

type WsHandler struct {
	hubServices input.HubService
}

func NewWsHandler(hubService input.HubService) *WsHandler {
	return &WsHandler{hubServices: hubService}
}

func (h *WsHandler) Ws(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrade:", err)
		return
	}
	h.hubServices.HandleConnection(&WsConn{conn})
}
