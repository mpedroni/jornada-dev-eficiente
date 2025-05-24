package author

import (
	"context"
	"database/sql"
	"desafiocdc/internal/author/domain"
)

type SqliteAuthorRepository struct {
	db *sql.DB
}

func (ar SqliteAuthorRepository) Save(ctx context.Context, a *domain.Author) error {
	result, err := ar.db.ExecContext(
		ctx,
		"INSERT INTO authors (name, email, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		a.Name(), a.Email(), a.Description(), a.CreatedAt(), a.UpdatedAt())
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	a.SetID(int(id))

	return nil
}
