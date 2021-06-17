package domain

import (
	"time"
)

type Candidate struct {
	Id      string
	Name    string
	Phone   string
	Email   string
	Address string
	Status  Status
}

type StatusEnum int

const (
	New StatusEnum = iota
	Offer
	Hire
	Decline
)

type Status struct {
	Type      StatusEnum
	StartedAt time.Time
	EndAt     *time.Time
}

type CandidateRepository interface {
	GetAll() (map[string]Candidate, error)
	GetById(id string) (Candidate, error)
	Delete(order *Candidate) error
	Store(order *Candidate) error
}
