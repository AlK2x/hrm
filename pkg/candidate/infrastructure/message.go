package infrastructure

import (
	"github.com/streadway/amqp"
	"hrm/pkg/candidate/domain"
)

type eventDispatcher struct {
	conn      *amqp.Connection
	ch        *amqp.Channel
	queue     amqp.Queue
	queueName string
}

type EventDispatcherConfig struct {
	Uri       string
	QueueName string
}

func NewReader(config EventDispatcherConfig) (domain.EventReader, error) {
	return createEventDispatcher(config)
}

func NewWriter(config EventDispatcherConfig) (domain.EventWriter, error) {
	return createEventDispatcher(config)
}

func createEventDispatcher(c EventDispatcherConfig) (*eventDispatcher, error) {
	conn, err := amqp.Dial(c.Uri)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(c.QueueName, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &eventDispatcher{
		conn:      conn,
		ch:        ch,
		queue:     q,
		queueName: c.QueueName,
	}, nil
}

func (e *eventDispatcher) Write(event domain.Event) error {
	return e.ch.Publish("", e.queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(event.Message),
	})
}

func (e *eventDispatcher) Read() (<-chan domain.Event, error) {
	ch := make(chan domain.Event)
	msgs, err := e.ch.Consume(e.queueName, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	go func() {
		for msg := range msgs {
			ch <- domain.Event{Message: string(msg.Body)}
		}
	}()

	return ch, nil
}
