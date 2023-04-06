package internal

import (
	"context"

	"go.uber.org/zap"

	"github.com/igomonov88/ms-users/internal/config"
	"github.com/igomonov88/ms-users/internal/db"
	"github.com/igomonov88/ms-users/internal/ms"
	"github.com/igomonov88/ms-users/internal/operation"
	"github.com/igomonov88/ms-users/internal/storage"
)

func Main() {
	cfg := config.Must(config.Read())
	msServer := newMSServer(cfg)
	msServer.Serve(context.Background())
}

func newMSServer(cfg config.Config) *ms.Server {
	logger := zap.NewExample()
	dbx := db.Must(db.New(cfg.DB))
	strg := storage.NewService(dbx)
	op := operation.NewService(strg)
	return ms.NewServer(cfg.MS, logger.Sugar(), op)
}
