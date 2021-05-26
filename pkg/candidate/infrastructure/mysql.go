package infrastructure

import "database/sql"

type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) sql.Row
	Commit() error
	Rollback() error
}
