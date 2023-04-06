package ms

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	httpServer "github.com/go-micro/plugins/v4/server/http"
	userspb "github.com/igomonov88/crea-genproto-go/creaapis/ms-users/v1"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"

	"github.com/igomonov88/ms-users/internal/http/api/handlers"
)

func (srv *Server) Serve(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	hServer := httpServer.NewServer(
		server.Name(srv.cfg.Name),
		server.Address(srv.cfg.HTTPAddress),
		server.Version(srv.cfg.Version),
		server.Context(ctx),
		server.Registry(registry.NewRegistry()),
	)

	router := handlers.Handler(srv.logger, srv.health)
	handler := hServer.NewHandler(router)
	if err := hServer.Handle(handler); err != nil {
		srv.logger.Fatalf("Server fail to handle: %v", err)
	}

	if err := hServer.Start(); err != nil {
		srv.logger.Fatalf("Server fail to start http server: %v", err)
	}

	service := micro.NewService(
		micro.Context(ctx),
		micro.Name(srv.cfg.Name),
		micro.Version(srv.cfg.Version),
		micro.Registry(registry.NewRegistry()),
		micro.Server(server.NewServer(server.Address(srv.cfg.Address))),
	)

	service.Init()

	userspb.RegisterUserServiceHandler(service.Server(), srv.users)
	// Handle signals gracefully
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		srv.logger.Infof("Server got signal %v, stopping", sig)
		cancel()
	}()
	{
		s := micro.NewService()
		client := userspb.NewUserService(srv.cfg.Address, s.Client())

		// call an endpoint on the service
		rsp, err := client.CreateUser(context.Background(), &userspb.User{
			Name: "John",
		})
		if err != nil {
			srv.logger.Fatalf("Server fail to call CreateUser: %v", err)
			return
		}

		// print the response
		srv.logger.Info("Response: ", rsp)
	}

	if err := service.Run(); err != nil {
		srv.logger.Fatalf("Server fail to serve: %v", err)
	}
}
