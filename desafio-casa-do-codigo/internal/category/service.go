package category

import (
	"context"
	"database/sql"
	"errors"
)

type CategoryService struct {
	db *sql.DB
}

func NewCategoryService(db *sql.DB) *CategoryService {
	return &CategoryService{
		db: db,
	}
}

func (s *CategoryService) Create(ctx context.Context, category *Category) error {
	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM categories WHERE name = ?)", category.Name).Scan(&exists)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if exists {
		return ErrCategoryAlreadyExists
	}

	result, err := s.db.ExecContext(ctx, "INSERT INTO categories (name, created_at, updated_at) VALUES (?, ?, ?)", category.Name, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		return err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	category.ID = int(lastInsertID)

	return nil
}
