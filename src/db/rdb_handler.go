package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type RDBHandler interface {
	Query(query string, args ...interface{}) (*sqlx.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
}