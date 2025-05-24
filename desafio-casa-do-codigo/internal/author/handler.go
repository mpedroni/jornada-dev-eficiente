package author

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type authorHandler struct {
	repo AuthorRepository
}

func (h *authorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var req CreateAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	author, err := req.toModel()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	if err := h.repo.Save(r.Context(), &author); err != nil {
		log.Default().Println(fmt.Errorf("saving author: %w", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	body := map[string]any{
		"id":          author.id,
		"name":        author.name,
		"email":       author.email,
		"description": author.description,
		"createdAt":   author.createdAt.Format(time.RFC3339),
	}

	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Default().Println(fmt.Errorf("encoding response body: %w", err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
