package category

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Setup(r *chi.Mux, db *sql.DB) {
	svc := CategoryService{db: db}
	handler := categoryHandler{svc: svc}

	registerRoutes(r, handler)
}

func registerRoutes(r *chi.Mux, h categoryHandler) {
	r.Route("/categories", func(r chi.Router) {
		r.Post("/", h.CreateCategory)
	})
}
