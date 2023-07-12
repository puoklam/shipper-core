package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/puoklam/shipper-core/http/route/users"
)

type RoutesHandler interface {
	SetupRoutes(r *chi.Mux)
}

func New() []RoutesHandler {
	return []RoutesHandler{
		&users.Handler{},
	}
}
