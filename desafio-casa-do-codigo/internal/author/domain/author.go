package domain

import (
	"errors"
	"time"
)

type Author struct {
	id          int
	name        string
	email       string
	description string
	createdAt   time.Time
	updatedAt   time.Time
}

func NewAuthor(name, email, description string) (Author, error) {
	now := time.Now().UTC()
	a := Author{
		id:          0,
		name:        name,
		email:       email,
		description: description,
		createdAt:   now,
		updatedAt:   now,
	}

	if err := a.selfValidate(); err != nil {
		return Author{}, err
	}

	return a, nil
}

type RestoreAuthorParams struct {
	ID          int
	Name        string
	Email       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func RestoreAuthor(p RestoreAuthorParams) (Author, error) {
	a := Author{
		id:          p.ID,
		name:        p.Name,
		email:       p.Email,
		description: p.Description,
		createdAt:   p.CreatedAt,
		updatedAt:   p.UpdatedAt,
	}

	if err := a.selfValidate(); err != nil {
		return Author{}, err
	}

	return a, nil
}

func (a Author) selfValidate() error {
	if a.Name() == "" {
		return errors.New("name is required")
	}

	if a.Email() == "" {
		return errors.New("email is required")
	}

	if a.Description() == "" {
		return errors.New("description is required")
	}

	if len(a.Description()) > 400 {
		return errors.New("description must be less than 400 characters")
	}

	return nil
}

func (a Author) ID() int {
	return a.id
}

// maybe the repository implementation should be at the same Author's level to avoid needing such SetID method
// but idk if i like the idea... also, in a self-generation id strategy (i.e uuid) such method wouldn't be necessary
func (a *Author) SetID(id int) {
	a.id = id
}

func (a Author) IsPersisted() bool {
	return a.id != 0
}

func (a Author) Name() string {
	return a.name
}

func (a Author) Email() string {
	return a.email
}

func (a Author) Description() string {
	return a.description
}

func (a Author) CreatedAt() time.Time {
	return a.createdAt
}

func (a Author) UpdatedAt() time.Time {
	return a.updatedAt
}
