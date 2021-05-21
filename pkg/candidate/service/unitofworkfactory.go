package service

type UnitOfWorkFactory interface {
	NewUnitOfWork() (UnitOfWork, error)
}
