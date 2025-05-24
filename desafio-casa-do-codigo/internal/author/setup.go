package author

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Setup(r *chi.Mux, db *sql.DB) {
	repo := AuthorRepository{db: db}
	handler := authorHandler{repo: repo}

	registerRoutes(r, handler)
}

func registerRoutes(r *chi.Mux, h authorHandler) {
	r.Route("/authors", func(r chi.Router) {
		r.Post("/", h.CreateAuthor)
	})
}
