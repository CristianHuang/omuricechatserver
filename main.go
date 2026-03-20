package main

import (
	"fmt"
	"net/http"

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

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Error upgrade:", err)
			return
		}
		defer conn.Close()

		fmt.Println("¡Celular conectado!")

		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Cliente desconectado")
				break
			}

			fmt.Printf("Recibido: %s\n", string(p))

			response := []byte("Servidor dice: Recibí: " + string(p))
			if err := conn.WriteMessage(messageType, response); err != nil {
				break
			}
		}
	})

	r.Run(":8080")
}
