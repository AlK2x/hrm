package model

import (
	"github.com/google/uuid"
	"time"
)

type CandidateOption func(*Candidate)

func WithName(name string) CandidateOption {
	return func(c *Candidate) {
		c.Name = name
	}
}

func WithPhone(phone string) CandidateOption {
	return func(c *Candidate) {
		c.Phone = phone
	}
}

func WithAddress(address string) CandidateOption {
	return func(c *Candidate) {
		c.Address = address
	}
}

type CandidateService struct {
	repository CandidateRepository
}

func (s *CandidateService) Register(options ...CandidateOption) (*Candidate, error) {
	c := &Candidate{Id: uuid.New().String()}

	for _, opt := range options {
		opt(c)
	}

	c.Status = Status{
		Type:      New,
		StartedAt: time.Time{},
	}

	return c, nil
}

func (s *CandidateService) MakeOffer(candidateId string) error {
	c, err := s.repository.GetById(candidateId)
	if err != nil {
		return err
	}

	c.Status.Type = Offer
	c.Status.StartedAt = time.Time{}

	return nil
}

func (s *CandidateService) Hire(candidateId string) error {
	c, err := s.repository.GetById(candidateId)
	if err != nil {
		return err
	}

	c.Status.Type = Hire
	c.Status.StartedAt = time.Time{}

	return nil
}

func (s *CandidateService) Decline(candidateId string) error {
	c, err := s.repository.GetById(candidateId)
	if err != nil {
		return err
	}

	c.Status.Type = Decline
	c.Status.StartedAt = time.Time{}

	return nil
}
