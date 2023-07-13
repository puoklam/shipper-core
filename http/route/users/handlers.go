package users

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/puoklam/shipper-core/db/models"
	htp "github.com/puoklam/shipper-core/http"
	"github.com/puoklam/shipper-core/http/validation"
	"github.com/volatiletech/sqlboiler/v4/boil"
	// . "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Handler struct {
}

func (h *Handler) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.Users().AllG(r.Context())

	if err != nil {
		htp.WriteError(w, http.StatusInternalServerError)
		return
	}

	if htp.WriteJson(w, users) == nil {
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := models.FindUserG(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			htp.WriteError(w, http.StatusNotFound)
		} else {
			htp.WriteError(w, http.StatusInternalServerError)
		}
		return
	}

	if htp.WriteJson(w, user) == nil {
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		htp.WriteError(w, http.StatusBadRequest)
	}

	vb := validation.NewUser().From(user)
	if err = validation.Vali.Struct(vb); err != nil {
		log.Println(err)
		htp.WriteError(w, http.StatusBadRequest)
	}
	vb.Done()

	err = user.InsertG(r.Context(), boil.Infer())
	if err != nil {
		log.Println(err)
		htp.WriteError(w, http.StatusInternalServerError)
		return
	}

	if htp.WriteJson(w, user) == nil {
		w.WriteHeader(http.StatusOK)
	}
}

func (h *Handler) SetupRoutes(r *chi.Mux) {
	r.Route("/users", func(r chi.Router) {
		// r.Use(middleware.Authenticator)
		r.Get("/", h.listUsers)
		r.Get("/{id}", h.getUser)
		r.Post("/", h.createUser)
	})
}
