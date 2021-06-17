package domain

type Message struct {
	Msg string
}

type MessageRepository interface {
	Save(msg Message) error
}
