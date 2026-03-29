package domain

type Message struct {
	SenderID string `json:"id"`
	Message  []byte `json:"message"`
}
type MessageSend struct {
	SenderID string `json:"id"`
	Message  []byte `json:"message"`
	SentAt   string `json:"sentAt"`
}
