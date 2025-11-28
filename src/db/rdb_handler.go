package db

import (
	"github.com/jmoiron/sqlx"
)

type RDBHandler interface {
	sqlx.ExtContext
}
