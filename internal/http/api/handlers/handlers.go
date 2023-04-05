package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"

	"github.com/igomonov88/ms-users/internal/usecase/health"
)

func Handler(log *zap.SugaredLogger, healthService *health.Service) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Route("/v1", func(r chi.Router) {
		healthHandler := NewHealthHandler(log, healthService)
		healthHandler.RegisterRoutes(r)
	})

	return router
}
