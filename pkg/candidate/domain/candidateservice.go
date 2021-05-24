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

func WithEmail(email string) CandidateOption {
	return func(c *Candidate) {
		c.Email = email
	}
}

type CandidateService struct {
	unitOfWorkFactory service.UnitOfWorkFactory
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

	unit, err := s.unitOfWorkFactory.NewUnitOfWork()
	if err != nil {
		return nil, err
	}
	err = unit.CandidateRepository().Add(c)
	if err != nil {
		return nil, err
	}
	unit.Complete(&err)
	return c, err
}

func (s *CandidateService) MakeOffer(candidateId string) error {
	unit, err := s.unitOfWorkFactory.NewUnitOfWork()
	if err != nil {
		return err
	}
	repo := unit.CandidateRepository()
	c, err := repo.GetById(candidateId)
	if err != nil {
		return err
	}

	c.Status.Type = Offer
	c.Status.StartedAt = time.Time{}

	err = repo.Update(c)
	unit.Complete(&err)
	return err
}

func (s *CandidateService) Hire(candidateId string) error {
	unit, err := s.unitOfWorkFactory.NewUnitOfWork()
	if err != nil {
		return err
	}
	candidateRepo := unit.CandidateRepository()
	c, err := candidateRepo.GetById(candidateId)
	if err != nil {
		return err
	}

	c.Status.Type = Hire
	c.Status.StartedAt = time.Time{}

	messageService := unit.MessageService()
	err = messageService.Send(Message{Msg: ""})
	if err != nil {
		return err
	}

	err = candidateRepo.Update(c)
	if err != nil {
		return err
	}

	unit.Complete(&err)
	return err
}

func (s *CandidateService) Decline(candidateId string) error {
	unit, err := s.unitOfWorkFactory.NewUnitOfWork()
	if err != nil {
		return err
	}
	repo := unit.CandidateRepository()
	c, err := repo.GetById(candidateId)
	if err != nil {
		return err
	}

	c.Status.Type = Decline
	c.Status.StartedAt = time.Time{}
	err = repo.Update(c)
	unit.Complete(&err)
	return err
}

func (s *CandidateService) validate(c *Candidate) error {
	if c.Name == "" {

	}
	return nil
}
