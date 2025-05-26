package author

import (
	"context"
	"database/sql"
	"desafiocdc/internal/author/domain"
	"errors"
	"time"
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

func (ar SqliteAuthorRepository) FindByEmail(ctx context.Context, email string) (domain.Author, error) {
	var id int64
	var name string
	var emailTarget string
	var description string
	var createdAt time.Time
	var updatedAt time.Time
	err := ar.db.
		QueryRow("SELECT id, name, email, description, created_at, updated_at FROM authors WHERE email = ?", email).
		Scan(&id, &name, &emailTarget, &description, &createdAt, &updatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Author{}, domain.ErrAuthorNotFound
		}
		return domain.Author{}, err
	}

	return domain.RestoreAuthor(domain.RestoreAuthorParams{
		ID:          int(id),
		Name:        name,
		Email:       email,
		Description: description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	})
}
