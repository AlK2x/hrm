package infrastructure

import (
	"hrm/pkg/candidate/domain"
)

func CreateRepository(tx Transaction) domain.CandidateRepository {
	return &MysqlCandidateRepository{tx: tx}
}

type MysqlCandidateRepository struct {
	tx Transaction
}

func (m *MysqlCandidateRepository) GetAll() (map[string]domain.Candidate, error) {
	panic("implement me")
}

func (m *MysqlCandidateRepository) GetById(id string) (*domain.Candidate, error) {
	panic("implement me")
}

func (m *MysqlCandidateRepository) Add(order *domain.Candidate) error {
	panic("implement me")
}

func (m *MysqlCandidateRepository) Delete(order *domain.Candidate) error {
	panic("implement me")
}

func (m *MysqlCandidateRepository) Update(order *domain.Candidate) error {
	panic("implement me")
}
