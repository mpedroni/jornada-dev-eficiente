package domain

import (
	"context"
	"errors"
)

var (
	ErrAuthorNotFound = errors.New("author not found")
)

type AuthorRepository interface {
	Save(ctx context.Context, a *Author) error
	FindByEmail(ctx context.Context, email string) (Author, error)
}
