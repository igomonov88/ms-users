package ms

import (
	"go.uber.org/zap"

	"github.com/igomonov88/ms-users/internal/operation"
	"github.com/igomonov88/ms-users/internal/usecase/health"
)

type Server struct {
	cfg    Config
	logger *zap.SugaredLogger
	health *health.Service
}

func NewServer(cfg Config, logger *zap.SugaredLogger, op *operation.Service) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		health: health.NewService(op, logger),
	}
}
