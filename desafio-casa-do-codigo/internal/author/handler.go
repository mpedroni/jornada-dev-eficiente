package author

import (
	"desafiocdc/internal/author/domain"
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
		ResponseError(w, NewHttpError(http.StatusBadRequest, err))
		return
	}

	author, err := req.toModel()
	if err != nil {
		BadRequest(w, err)
		return
	}

	possibleAuthor, err := h.repo.FindByEmail(r.Context(), req.Email)
	if err != nil && !errors.Is(err, domain.ErrAuthorNotFound) {
		InternalServerError(w, err)
		return
	}

	if possibleAuthor.IsPersisted() {
		Conflict(w, fmt.Errorf("there is already an author registered with the email %s", req.Email))
		return
	}

	if err := h.repo.Save(r.Context(), &author); err != nil {
		log.Default().Println(fmt.Errorf("saving author: %w", err))
		InternalServerError(w, err)
		return
	}

	if err := json.NewEncoder(w).Encode(CreateAuthorResponseFrom(author)); err != nil {
		log.Default().Println(fmt.Errorf("encoding response body: %w", err))
		InternalServerError(w, err)
		return
	}
}
