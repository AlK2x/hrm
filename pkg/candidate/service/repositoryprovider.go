package service

import "hrm/pkg/candidate/domain"

type RepositoryProvider interface {
	MessageService() domain.MessageService
	CandidateRepository() domain.CandidateRepository
}

type UnitOfWork interface {
	RepositoryProvider
	Complete(e *error)
}
