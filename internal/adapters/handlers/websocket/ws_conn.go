package websocket

import (
	"time"

	"github.com/gorilla/websocket"
)

type WsConn struct {
	conn *websocket.Conn
}

func (w *WsConn) ReadMessage() ([]byte, error) {
	_, msg, err := w.conn.ReadMessage()
	return msg, err
}

func (w *WsConn) WriteMessage(number int, msg []byte) error {
	return w.conn.WriteMessage(websocket.TextMessage, msg)
}

func (w *WsConn) Close() error {
	return w.conn.Close()
}

func (w *WsConn) SetWriteDeadline(t time.Time) {
	w.conn.SetWriteDeadline(t)
}
