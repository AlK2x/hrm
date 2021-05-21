package domain

import (
	"github.com/google/uuid"
	"hrm/pkg/candidate/service"
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
	repository     CandidateRepository
	messageService MessageService
	unitOfWork     service.UnitOfWork
}

func (s *CandidateService) Register(options ...CandidateOption) (*Candidate, error) {
	c := &Candidate{Id: uuid.New().String()}

	for _, opt := range options {
		opt(c)
	}

	err := s.validate(c)
	if err != nil {
		return nil, err
	}

	c.Status = Status{
		Type:      New,
		StartedAt: time.Time{},
	}

	err = s.repository.Add(c)
	if err != nil {
		return nil, err
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

	err = s.unitOfWork.MessageService().Send(Message{Msg: ""})
	if err != nil {
		return err
	}

	err = s.unitOfWork.CandidateRepository().Update(c)
	if err != nil {
		return err
	}

	return s.unitOfWork.Complete(err)
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

func (s *CandidateService) validate(c *Candidate) error {
	if c.Name == "" {

	}
	return nil
}
