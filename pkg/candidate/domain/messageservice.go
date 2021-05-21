package domain

type Message struct {
	Msg string
}

type MessageService interface {
	Send(msg Message) error
}
