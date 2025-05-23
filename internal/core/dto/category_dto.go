package dto

import (
	"finhub-go/internal/core/domain"

	"github.com/google/uuid"
)

type CategoryRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Color       *string `json:"color"`
}

type CategoryResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Color       *string   `json:"color"`
}

func NewCategoryResponse(id uuid.UUID, name string, description, color *string) *CategoryResponse {
	return &CategoryResponse{
		ID:          id,
		Name:        name,
		Description: description,
		Color:       color,
	}
}

func (r *CategoryRequest) ToDomain() (*domain.Category, error) {
	return domain.NewCategory(r.Name, r.Description, r.Color)

}
