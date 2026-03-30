package websocket

import (
	"fmt"
	"net/http"

	"github.com/cristianhuang/omuricechatserver/internal/core/domain"
	"github.com/cristianhuang/omuricechatserver/internal/core/services"
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
	hubStruct    *domain.Hub
	roomServices services.RoomServiceInterface
}

func NewWsHandler(roomService services.RoomServiceInterface) *WsHandler {
	return &WsHandler{
		hubStruct:    domain.NewHub(),
		roomServices: roomService,
	}
}

func (w *WsHandler) HandleRoom(c *gin.Context) {
	id := c.Param("id")

	room, ok := w.hubStruct.GetRoom(id)

	if !ok {
		room = w.hubStruct.CreateRoom(id)
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrade:", err)
		return
	}
	w.roomServices.HandleConnection(&WsConn{conn}, room)

}
