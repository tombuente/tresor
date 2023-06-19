package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/tombuente/tresor/service/health"
)

type healthHandler struct {
	service health.Service
}

func NewHealthRouter(service health.Service) chi.Router {
	handler := healthHandler{
		service: service,
	}

	router := chi.NewRouter()
	router.Get("/", handler.getHealth)

	return router
}

func (handler healthHandler) getHealth(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, handler.service.GetHealth(r.Context()))
}
