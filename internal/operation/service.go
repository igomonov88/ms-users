package operation

import (
	"github.com/igomonov88/ms-users/internal/operation/health"
	"github.com/igomonov88/ms-users/internal/storage"
)

type Service struct {
	Health *health.Service
}

func NewService(storage *storage.Service) *Service {
	return &Service{
		Health: health.NewService(storage),
	}
}
