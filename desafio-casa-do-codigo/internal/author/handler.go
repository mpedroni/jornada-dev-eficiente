package author

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type authorHandler struct {
	repo AuthorRepository
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
