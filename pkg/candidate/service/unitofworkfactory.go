package service

import "hrm/pkg/candidate/domain"

type UnitOfWorkFactory interface {
	NewUnitOfWork() (domain.CandidateUnitOfWork, error)
}
