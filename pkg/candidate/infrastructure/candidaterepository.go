package infrastructure

import (
	"database/sql"
	"errors"
	"hrm/pkg/candidate/domain"
)

type MysqlCandidateRepository struct {
	tx *sql.Tx
}

func (m *MysqlCandidateRepository) Store(candidate *domain.Candidate) error {
	c, err := m.GetById(candidate.Id)
	if err == domain.ErrorCandidateNotExist {
		query := `INSERT INTO  candidate (id, name, email, address, phone)
			VALUES(?,?,?,?,?,?,NOW())`
		statusQuery := `INSERT INTO status (candidate_id, type, created_at) VALUES (?, ?, NOW())`
		_, err = m.tx.Exec(query, candidate.Id, candidate.Email, candidate.Name, candidate.Address, candidate.Phone)
		if err != nil {
			return nil
		}
		_, err = m.tx.Exec(statusQuery, candidate.Id, candidate.Status.Type)
		if err != nil {
			return nil
		}
	} else if err != nil {
		return err
	} else {
		query := `UPDATE candidate c SET c.name = ?, c.email = ?, c.address = ?, c.phone = ? WHERE c.id = ?`
		newStatusQuery := `INSERT INTO status (candidate_id, type, created_at) VALUES (?, ?, ?)`
		updateOldStatus := `UPDATE status SET s.deleted_at = NOW() WHERE s.candidate_id = ? AND s.deleted_at IS NULL`
		_, err = m.tx.Exec(query, candidate.Name, candidate.Email, candidate.Address, candidate.Phone, c.Id)
		if err != nil {
			return nil
		}
		_, err = m.tx.Exec(newStatusQuery, candidate.Id, candidate.Status.StartedAt)
		if err != nil {
			return nil
		}
		_, err = m.tx.Exec(updateOldStatus, candidate.Id)
		if err != nil {
			return nil
		}
	}

	return nil
}

func (m *MysqlCandidateRepository) GetAll() (map[string]domain.Candidate, error) {
	query := `SELECT c.id, c.name, c.email, c.address, c.phone, s.type, s.created_at FROM candidate c INNER JOIN status s ON s.candidate_id = c.id AND s.deleted_at IS NULL`
	rows, err := m.tx.Query(query)
	if err != nil {
		return nil, nil
	}

	defer rows.Close()

	var candidateId string
	var name string
	var address string
	var phone string
	var status int
	var email string
	var createdAt sql.NullTime

	candidates := make(map[string]domain.Candidate)
	for rows.Next() {
		err := rows.Scan(&candidateId, &name, &email, &address, &phone, &status, &createdAt)
		if err != nil {
			return candidates, err
		}

		order, ok := candidates[candidateId]
		if !ok {
			order = domain.Candidate{
				Id:      candidateId,
				Name:    name,
				Phone:   phone,
				Email:   email,
				Address: address,
				Status: domain.Status{
					Type:      domain.StatusEnum(status),
					StartedAt: createdAt.Time,
				},
			}
		}

		candidates[candidateId] = order
	}

	return candidates, nil
}

func (m *MysqlCandidateRepository) GetById(id string) (domain.Candidate, error) {
	query := `SELECT c.id, c.name, c.address, c.phone, s.type, s.created_at 
		FROM candidate c INNER JOIN status s ON s.candidate_id = c.id AND s.deleted_at IS NULL WHERE c.id = ?`

	row := m.tx.QueryRow(query, id)
	var candidate domain.Candidate
	err := row.Scan(&candidate)
	if errors.Is(err, sql.ErrNoRows) {
		return candidate, domain.ErrorCandidateNotExist
	}
	return candidate, nil
}

func (m *MysqlCandidateRepository) Delete(candidate *domain.Candidate) error {
	_, err := m.tx.Exec("DELETE FROM `candidate` WHERE id = ?", candidate.Id)
	if err != nil {
		return err
	}

	return nil
}
