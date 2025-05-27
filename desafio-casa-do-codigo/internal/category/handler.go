package category

import (
	"desafiocdc/pkg/http/httperror"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type categoryHandler struct {
	svc CategoryService
}

func (h *categoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	category, err := req.toModel()
	if err != nil {
		httperror.BadRequest(w, err)
		return
	}
	if err := h.svc.Create(r.Context(), &category); err != nil {
		if err == ErrCategoryAlreadyExists {
			httperror.Conflict(w, fmt.Errorf("there already exists a category with name %s", category.Name))
			return
		}
		log.Default().Println(fmt.Errorf("creating category: %w", err))
		httperror.InternalServerError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(CreateCategoryResponseFrom(category)); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
