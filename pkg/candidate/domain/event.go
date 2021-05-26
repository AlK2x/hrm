package domain

import "time"

type EventType string

const (
	CandidateHiredEventType = EventType("candidate_hired")
)

type Event interface {
	GetType() EventType
}

type CandidateHired struct {
	CandidateId string
	Date        time.Time
}

func (c CandidateHired) GetType() EventType {
	return CandidateHiredEventType
}

type EventDispatcher interface {
	Dispatch(Event) error
}

type EventWriter interface {
	Write(Event) error
}

type EventConsumer interface {
	Read() (<-chan Event, error)
}
