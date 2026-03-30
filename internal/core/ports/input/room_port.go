package input

import (
	"time"
)

type Connection interface {
	ReadMessage() ([]byte, error)
	WriteMessage(int, []byte) error
	SetWriteDeadline(time.Time)
	Close() error
}
