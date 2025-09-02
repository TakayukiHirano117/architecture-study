package db

import (
	// "database/sql"

	"github.com/jmoiron/sqlx"
)

type RDBHandler interface {
	sqlx.ExtContext
}