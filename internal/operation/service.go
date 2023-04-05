package operation

import "github.com/igomonov88/ms-users/internal/operation/health"

type Service struct {
	Health *health.Service
}

func NewService() *Service {
	return &Service{
		Health: health.NewService(),
	}
}
