package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.uber.org/zap"

	"github.com/igomonov88/ms-users/internal/http/api/contract"
	"github.com/igomonov88/ms-users/internal/usecase/health"
)

type HealthHandler struct {
	logger  *zap.SugaredLogger
	service *health.Service
}

func NewHealthHandler(logger *zap.SugaredLogger, service *health.Service) HealthHandler {
	return HealthHandler{
		logger:  logger,
		service: service,
	}
}

func (h HealthHandler) RegisterRoutes(router chi.Router) {
	router.Get("/health", h.checkHealth)
}

func (h HealthHandler) checkHealth(w http.ResponseWriter, r *http.Request) {
	if err := h.service.CheckHealth(r.Context()); err != nil {
		h.logger.Errorw("health check failed", "error", err)
		render.Render(w, r, contract.ErrorInternal)
		return
	}

	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, nil)
}
