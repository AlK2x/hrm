package domain

type Event struct {
	Message string
}

type CandidateHired struct {
	Event
}

type EventWriter interface {
	Write(Event) error
}

type EventReader interface {
	Read() (<-chan Event, error)
}
