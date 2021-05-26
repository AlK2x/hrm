package domain

type UnitOfWorkFactory interface {
	NewUnitOfWork() (CandidateUnitOfWork, error)
}
