package infrastructure

import "database/sql"

type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Commit() error
	Rollback() error
}
