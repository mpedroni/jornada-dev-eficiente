package author

import (
	"desafiocdc/internal/author/domain"
	"encoding/json"
	"net/http"
	"time"
)

type HttpError struct {
	Err        error `json:"-"`
	StatusCode int   `json:"-"`

	StatusText string `json:"status"`
	ErrorText  string `json:"error"`
}

func NewHttpError(status int, err error) HttpError {
	return HttpError{
		Err:        err,
		StatusCode: status,

		StatusText: http.StatusText(status),
		ErrorText:  err.Error(),
	}
}

func ResponseError(w http.ResponseWriter, err HttpError) {
	w.WriteHeader(err.StatusCode)
	_ = json.NewEncoder(w).Encode(err)
}

func BadRequest(w http.ResponseWriter, err error) {
	ResponseError(w, NewHttpError(http.StatusBadRequest, err))
}

func InternalServerError(w http.ResponseWriter, err error) {
	ResponseError(w, NewHttpError(http.StatusInternalServerError, err))
}

func Conflict(w http.ResponseWriter, err error) {
	ResponseError(w, NewHttpError(http.StatusConflict, err))
}

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
