package infrastructure

import (
	"database/sql"
	"hrm/pkg/candidate/domain"
)

type unitOfWorkFactory struct {
	client sql.DB
}

type unitOfWork struct {
	tx Transaction
}

func (u *unitOfWork) CandidateRepository() domain.CandidateRepository {
	return CreateRepository(u.tx)
}

func (u *unitOfWork) Complete(err error) {

}

func (u *unitOfWorkFactory) NewUnitOfWork() (*unitOfWork, error) {
	tx, err := u.client.Begin()
	if err != nil {
		return nil, err
	}

	return &unitOfWork{tx: tx}, nil
}
