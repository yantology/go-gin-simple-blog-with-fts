package models

//message model
type Message struct {
	Message string `json:"message"`
}

func NewMessage(message string) *Message {
	return &Message{
		Message: message,
	}
}
