package db

import (
	"database/sql"

	"github.com/pkg/errors"
)

const (
	ErrNotFoundCode            = "not found"
	ErrNoAffectedRowsCode      = "no affected rows"
	ErrDuplicateCode           = "duplicate"
	ErrForeignKeyViolationCode = "foreign key violation"
)

var (
	ErrNotFound            = errors.New(ErrNotFoundCode)
	ErrNoAffectedRows      = errors.New(ErrNoAffectedRowsCode)
	ErrDuplicate           = errors.New(ErrDuplicateCode)
	ErrForeignKeyViolation = errors.New(ErrForeignKeyViolationCode)
)

func WrapError(err error, message string) error {
	causeErr := errors.Cause(err)
	switch {
	case causeErr == sql.ErrNoRows:
		return errors.WithMessage(ErrNotFound, message)
	case IsDuplicateKey(causeErr):
		return errors.WithMessage(ErrDuplicate, message)
	case IsForeignKeyViolation(causeErr):
		return errors.WithMessage(ErrForeignKeyViolation, message)
	default:
		return errors.Wrap(err, message)
	}
}

func CheckAffectedRows(rs sql.Result) error {
	c, err := rs.RowsAffected()
	if err != nil {
		return WrapError(err, "failed to get affected rows")
	}

	if c == 0 {
		return ErrNoAffectedRows
	}

	return nil
}
