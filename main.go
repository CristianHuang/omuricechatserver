package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// ESTO ES LO QUE TE FALTA:
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Esto permite que tu celular se conecte sin errores de CORS
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		// 1. Upgrade de HTTP a WebSocket
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Error upgrade:", err)
			return
		}
		defer conn.Close()

		fmt.Println("¡Celular conectado!")

		for {
			// 2. Leer mensaje
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Cliente desconectado")
				break
			}

			fmt.Printf("Recibido: %s\n", string(p))

			// 3. Responder (Eco)
			response := []byte("Servidor dice: Recibí: " + string(p))
			if err := conn.WriteMessage(messageType, response); err != nil {
				break
			}
		}
	})

	r.Run(":8080")
}
