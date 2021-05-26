package infrastructure

import (
	"database/sql"
	"errors"
	"hrm/pkg/candidate/domain"
)

func CreateRepository(tx Transaction) domain.CandidateRepository {
	return &MysqlCandidateRepository{tx: tx}
}

type MysqlCandidateRepository struct {
	tx Transaction
}

func (m *MysqlCandidateRepository) Update(order *domain.Candidate) error {
	panic("implement me")
}

func (m *MysqlCandidateRepository) GetAll() (map[string]domain.Candidate, error) {
	panic("implement me")
}

func (m *MysqlCandidateRepository) GetById(id string) (domain.Candidate, error) {
	query := `SELECT c.id, c.name, c.address, c.phone, s.type, s.created_at 
		FROM candidate c INNER JOIN status s ON s.candidate_id = c.id AND s.deleted_at IS NULL WHERE c.id = ?`

	row := m.tx.QueryRow(query, id)
	var candidate domain.Candidate
	err := row.Scan(&candidate)
	if errors.Is(err, sql.ErrNoRows) {
		return candidate, domain.ErrorCandidateNotExists
	}
	return candidate, nil
}

func (m *MysqlCandidateRepository) Store(order *domain.Candidate) error {
	panic("implement me")
}

func (m *MysqlCandidateRepository) Delete(order *domain.Candidate) error {
	panic("implement me")
}
