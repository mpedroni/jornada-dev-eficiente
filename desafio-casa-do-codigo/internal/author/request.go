package author

type CreateAuthorRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

func (r CreateAuthorRequest) toModel() (Author, error) {
	return NewAuthor(r.Name, r.Email, r.Description)
}
