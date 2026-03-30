package domain

type Message struct {
	ID      string `json:"id"`
	Message []byte `json:"message"`
}
type MessageSend struct {
	ID      string `json:"id"`
	Message []byte `json:"message"`
	SentAt  string `json:"sentAt"`
}
