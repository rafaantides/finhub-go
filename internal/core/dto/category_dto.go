package dto

import (
	"finhub-go/internal/core/domain"

	"github.com/google/uuid"
)

type CategoryRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CategoryResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

func NewCategoryResponse(id uuid.UUID, name string, description *string) *CategoryResponse {
	return &CategoryResponse{
		ID:          id,
		Name:        name,
		Description: description,
	}
}

func (r *CategoryRequest) ToDomain() (*domain.Category, error) {
	return domain.NewCategory(r.Name, r.Description)

}
