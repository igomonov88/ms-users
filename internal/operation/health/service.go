package health

import (
	"context"

	"github.com/igomonov88/ms-users/internal/storage"
)

type Service struct {
	storage *storage.Service
}

func NewService(storage *storage.Service) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) HealthCheck(ctx context.Context) error {
	return s.storage.Health.Ping(ctx)
}
