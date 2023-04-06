package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/igomonov88/ms-users/internal/storage/health"
)

type Service struct {
	Health *health.Storage
}

func NewService(dbx *sqlx.DB) *Service {
	return &Service{
		Health: health.NewStorage(dbx),
	}
}
