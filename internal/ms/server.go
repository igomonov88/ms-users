package ms

import (
	"go.uber.org/zap"

	"github.com/igomonov88/ms-users/internal/operation"
	"github.com/igomonov88/ms-users/internal/usecase/health"
	"github.com/igomonov88/ms-users/internal/usecase/users"
)

type Server struct {
	cfg    Config
	logger *zap.SugaredLogger
	health *health.Service
	users  *users.Service
}

func NewServer(cfg Config, logger *zap.SugaredLogger, op *operation.Service) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
		health: health.NewService(op, logger),
		users:  users.NewService(op, logger),
	}
}
