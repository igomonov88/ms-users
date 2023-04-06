package db

import (
	"github.com/jackc/pgx"
)

const (
	pqUniqueViolation     = "23505"
	pqForeignKeyViolation = "23503"
)

func IsDuplicateKey(err error) bool {
	if err != nil {
		switch e := err.(type) {
		case pgx.PgError:
			return e.Code == pqUniqueViolation
		}
	}

	return false
}

func IsForeignKeyViolation(err error) bool {
	if err == nil {
		switch e := err.(type) {
		case pgx.PgError:
			return e.Code == pqForeignKeyViolation
		}
	}

	return false
}
