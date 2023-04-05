package ms

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	httpServer "github.com/go-micro/plugins/v4/server/http"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"

	"github.com/igomonov88/ms-users/internal/http/api/handlers"
)

func (srv *Server) Serve(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	server := httpServer.NewServer(
		server.Name(srv.cfg.Name),
		server.Address(srv.cfg.Address),
		server.Version(srv.cfg.Version),
		server.Context(ctx),
	)

	router := handlers.Handler(srv.logger, srv.health)
	handler := server.NewHandler(router)
	if err := server.Handle(handler); err != nil {
		srv.logger.Fatalf("Server fail to handle: %v", err)
	}

	service := micro.NewService(
		micro.Server(server),
		micro.Registry(registry.NewRegistry()),
	)

	service.Init()

	// Handle signals gracefully
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		srv.logger.Infof("Server got signal %v, stopping", sig)
		cancel()
	}()

	if err := service.Run(); err != nil {
		srv.logger.Fatalf("Server fail to serve: %v", err)
	}
}
