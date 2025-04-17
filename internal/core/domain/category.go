package domain

import (
	"finhub-go/internal/core/errors"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID
	Name        string
	Description *string
	CreatedAt   string
	UpdatedAt   string
}

func NewCategory(name string, description *string) (*Category, error) {

	if name == "" {
		return nil, errors.EmptyField("name")
	}

	return &Category{
		Name:        name,
		Description: description,
	}, nil
}
