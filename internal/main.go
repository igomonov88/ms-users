package internal

import (
	"context"

	"go.uber.org/zap"

	"github.com/igomonov88/ms-users/internal/config"
	"github.com/igomonov88/ms-users/internal/ms"
	"github.com/igomonov88/ms-users/internal/operation"
)

func Main() {
	cfg := config.Must(config.Read())
	msServer := newMSServer(cfg)
	msServer.Serve(context.Background())
}

func newMSServer(cfg config.Config) *ms.Server {
	logger := zap.NewExample()
	op := operation.NewService()
	return ms.NewServer(cfg.MS, logger.Sugar(), op)
}
