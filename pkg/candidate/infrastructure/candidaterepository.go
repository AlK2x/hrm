package infrastructure

import (
	"database/sql"
	"hrm/pkg/candidate/model"
)

func CreateRepository(db *sql.DB) model.CandidateRepository {
	return &MysqlCandidateRepository{db: db}
}

type MysqlCandidateRepository struct {
	db *sql.DB
}

func (m *MysqlCandidateRepository) GetAll() (map[string]model.Candidate, error) {
	panic("implement me")
}

func (m *MysqlCandidateRepository) GetById(id string) (*model.Candidate, error) {
	panic("implement me")
}

func (m *MysqlCandidateRepository) Add(order *model.Candidate) error {
	panic("implement me")
}

func (m *MysqlCandidateRepository) Delete(order *model.Candidate) error {
	panic("implement me")
}

func (m *MysqlCandidateRepository) Update(order *model.Candidate) error {
	panic("implement me")
}
