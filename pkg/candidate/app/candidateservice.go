package app

import (
	"github.com/google/uuid"
	"hrm/pkg/candidate/domain"
	"time"
)

type CandidateOption func(*domain.Candidate)

func WithName(name string) CandidateOption {
	return func(c *domain.Candidate) {
		c.Name = name
	}
}

func WithPhone(phone string) CandidateOption {
	return func(c *domain.Candidate) {
		c.Phone = phone
	}
}

func WithAddress(address string) CandidateOption {
	return func(c *domain.Candidate) {
		c.Address = address
	}
}

func WithEmail(email string) CandidateOption {
	return func(c *domain.Candidate) {
		c.Email = email
	}
}

type CandidateService struct {
	unitOfWorkFactory domain.UnitOfWorkFactory
}

func (s *CandidateService) Register(options ...CandidateOption) (*domain.Candidate, error) {
	c := &domain.Candidate{Id: uuid.New().String()}

	for _, opt := range options {
		opt(c)
	}

	err := s.validate(c)
	if err != nil {
		return nil, err
	}

	c.Status = domain.Status{
		Type:      domain.New,
		StartedAt: time.Time{},
	}

	unit, err := s.unitOfWorkFactory.NewUnitOfWork()
	if err != nil {
		return nil, err
	}
	err = unit.CandidateRepository().Update(c)
	if err != nil {
		return nil, err
	}
	defer unit.Complete(&err)
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

	c.Status.Type = domain.Offer
	c.Status.StartedAt = time.Time{}

	err = repo.Update(&c)

	defer unit.Complete(&err)

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

	c.Status.Type = domain.Hire
	c.Status.StartedAt = time.Time{}

	messageService := unit.MessageRepository()
	err = messageService.Save(domain.Message{Msg: ""})
	if err != nil {
		return err
	}

	err = candidateRepo.Update(&c)
	if err != nil {
		return err
	}

	defer unit.Complete(&err)
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

	c.Status.Type = domain.Decline
	c.Status.StartedAt = time.Time{}
	err = repo.Update(&c)
	defer unit.Complete(&err)
	return err
}

func (s *CandidateService) validate(c *domain.Candidate) error {
	if c.Name == "" {

	}
	return nil
}
