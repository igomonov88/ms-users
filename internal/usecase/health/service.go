package health

import (
	"context"

	"go.uber.org/zap"

	"github.com/igomonov88/ms-users/internal/operation"
)

type Service struct {
	op     *operation.Service
	logger *zap.SugaredLogger
}

func (s *Service) CheckHealth(ctx context.Context) error {
	return s.op.Health.CheckHealth(ctx)
}

func NewService(op *operation.Service, logger *zap.SugaredLogger) *Service {
	return &Service{
		op:     op,
		logger: logger,
	}
}
