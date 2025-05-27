package category

import (
	"errors"
	"time"
)

var (
	ErrCategoryAlreadyExists = errors.New("category already exists")
)

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCategory(name string) (Category, error) {
	if name == "" {
		return Category{}, errors.New("name is required")
	}

	now := time.Now().UTC()

	return Category{
		ID:        0,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
