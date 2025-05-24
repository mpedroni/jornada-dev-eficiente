package author

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
		id:          -1,
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
