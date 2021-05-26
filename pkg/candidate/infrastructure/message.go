package infrastructure

import (
	"encoding/json"
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

func NewReader(config EventDispatcherConfig) (domain.EventConsumer, error) {
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
	eventData, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return e.ch.Publish("", e.queue.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        eventData,
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
			var e domain.CandidateHired
			err = json.Unmarshal(msg.Body, &e)
			if err != nil {
				continue
			}
			ch <- e
		}
	}()

	return ch, nil
}

type mysqlMessageRepository struct {
	tx Transaction
}

func (d *mysqlMessageRepository) Save(msg domain.Message) error {
	return nil
}
