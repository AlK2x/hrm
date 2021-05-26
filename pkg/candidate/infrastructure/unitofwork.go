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

func (u *unitOfWork) MessageRepository() domain.MessageRepository {
	return &mysqlMessageRepository{tx: u.tx}
}

func (u *unitOfWork) CandidateRepository() domain.CandidateRepository {
	return CreateRepository(u.tx)
}

func (u *unitOfWork) Complete(err *error) {
	if err != nil {
		u.tx.Rollback()
	} else {
		err2 := u.tx.Commit()
		err = &err2
	}
}

func (u *unitOfWorkFactory) NewUnitOfWork() (domain.CandidateUnitOfWork, error) {
	tx, err := u.client.Begin()
	if err != nil {
		return nil, err
	}

	return &unitOfWork{tx: tx}, nil
}
