package domain

type RepositoryProvider interface {
	MessageRepository() MessageRepository
	CandidateRepository() CandidateRepository
}

type CandidateUnitOfWork interface {
	RepositoryProvider
	UnitOfWork
}

type UnitOfWork interface {
	Complete(e *error)
}
