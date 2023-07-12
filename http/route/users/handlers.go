package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/puoklam/shipper-core/db/models"
	htp "github.com/puoklam/shipper-core/http"
	"github.com/puoklam/shipper-core/middleware"
)

type Handler struct {
}

func (h *Handler) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.Users().AllG(r.Context())
	if err != nil {
		htp.WriteInternalErr(w)
		return
	}
	if htp.WriteJson(w, users) == nil {
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) SetupRoutes(r *chi.Mux) {
	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.Authenticator)
		r.Get("/", h.listUsers)
	})
}
