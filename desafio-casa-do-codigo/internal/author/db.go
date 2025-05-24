package author

import (
	"context"
	"database/sql"
)

type AuthorRepository struct {
	db *sql.DB
}

func (ar *AuthorRepository) Save(ctx context.Context, a *Author) error {
	result, err := ar.db.ExecContext(
		ctx,
		"INSERT INTO authors (name, email, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		a.name, a.email, a.description, a.createdAt, a.updatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	a.id = int(id)

	return nil
}
