package rdb

import (
	"context"
	"fmt"

	"github.com/TakayukiHirano117/architecture-study/config"
	"github.com/TakayukiHirano117/architecture-study/src/db"
	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
)

func NewConnection(c *config.DBConfig) (*sqlx.DB, error) {
	var conn *sqlx.DB
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)

	conn, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, errors.New("fail to connect to database")
	}

	return conn, nil
}

func ExecFromCtx(ctx context.Context) (db.RDBHandler, error) {
	val := ctx.Value(config.DBKey)

	if val == nil {
		return nil, errors.New("fail to get connection from context")
	}

	conn, ok := val.(db.RDBHandler)
	if !ok {
		return nil, errors.New("can't get context executor")
	}

	return conn, nil
}
