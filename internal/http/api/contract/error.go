package contract

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/go-chi/chi/middleware"
)

type ErrorResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	RequestID      string `json:"request_id"`
	StatusText     string `json:"status"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	e.RequestID = middleware.GetReqID(r.Context())
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var (
	ErrorInternal = &ErrorResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal server error.",
	}
	ErrorBadRequest = &ErrorResponse{
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Bad request.",
	}
)
