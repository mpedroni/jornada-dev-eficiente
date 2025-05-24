package domain

import "context"

type AuthorRepository interface {
	Save(ctx context.Context, a *Author) error
}
