package db

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func Must(db *sqlx.DB, err error) *sqlx.DB {
	if err != nil {
		panic(err)
	}

	return db
}

func New(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect(cfg.DriverName, cfg.DSN)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create db connection")
	}

	maxOpenConns := cfg.MaxOpenConns
	if maxOpenConns <= 0 {
		maxOpenConns = DefaultMaxOpenConns
	}

	maxIdleConns := cfg.MaxIdleConns
	if maxIdleConns <= 0 {
		maxIdleConns = DefaultMaxIdleConns
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	return db, nil
}
