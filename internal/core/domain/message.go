package domain

type Message struct {
	Body   string
	RoomID string
	From   string
}

type Typing struct {
	Typing bool
	From   string
}
