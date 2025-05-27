package category

import "time"

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func (r *CreateCategoryRequest) toModel() (Category, error) {
	return NewCategory(r.Name)
}

type CreateCategoryResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func CreateCategoryResponseFrom(category Category) CreateCategoryResponse {
	return CreateCategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format(time.RFC3339),
	}
}
