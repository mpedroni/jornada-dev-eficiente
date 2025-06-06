package author

import (
	"desafiocdc/internal/author/domain"
	"desafiocdc/pkg/http/httperror"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type authorHandler struct {
	repo domain.AuthorRepository
}

func (h *authorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var req CreateAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, fmt.Errorf("unable to decode request body: %w", err))
		return
	}

	author, err := req.toModel()
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}

	possibleAuthor, err := h.repo.FindByEmail(r.Context(), req.Email)
	if err != nil && !errors.Is(err, domain.ErrAuthorNotFound) {
		log.Default().Println(fmt.Errorf("finding author by email: %w", err))
		httperror.InternalServerError(w, err)
		return
	}

	if possibleAuthor.IsPersisted() {
		httperror.Conflict(w, fmt.Errorf("there already exists an author with email %s", req.Email))
		return
	}

	if err := h.repo.Save(r.Context(), &author); err != nil {
		log.Default().Println(fmt.Errorf("saving author: %w", err))
		httperror.InternalServerError(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(CreateAuthorResponseFrom(author)); err != nil {
		log.Default().Println(fmt.Errorf("encoding response body: %w", err))
		httperror.InternalServerError(w, err)
		return
	}
}
