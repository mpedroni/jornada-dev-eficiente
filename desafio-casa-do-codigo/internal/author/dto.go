package author

import (
	"desafiocdc/internal/author/domain"
	"time"
)

type CreateAuthorRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

func (r CreateAuthorRequest) toModel() (domain.Author, error) {
	return domain.NewAuthor(r.Name, r.Email, r.Description)
}

type CreateAuthorResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

func CreateAuthorResponseFrom(a domain.Author) CreateAuthorResponse {
	return CreateAuthorResponse{
		ID:          a.ID(),
		Name:        a.Name(),
		Email:       a.Email(),
		Description: a.Description(),
		CreatedAt:   a.CreatedAt().UTC().Format(time.RFC3339),
	}
}
