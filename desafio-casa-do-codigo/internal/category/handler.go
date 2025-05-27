package category

import (
	"desafiocdc/pkg/http/httperror"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
Here in category package we are using a different approach compared to the [internal/author] package.
Instead of using the handler as the service layer, we are actually using a separate service layer
to handle the business logic, but the service is dealing with the database directly (SQL directly).

Although I don't like that much having business logic mixed up with database concerns, it seems to be acceptable
in this context since the application is small and the business logic is simple.
Besides, it would't make much sense to test the service layer without the database anyway.

This approach allows us to keep the handler layer thin, focusing on request handling and response formatting. In addition of that,
we can always refactor the service layer later if we need to add more complex business logic or if we want to introduce a repository pattern without impacting the handler layer.
*/
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
