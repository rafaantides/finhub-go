package dto

import (
	"finhub-go/internal/core/domain"

	"github.com/google/uuid"
)

type PaymentStatusRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type PaymentStatusResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

func NewPaymentStatusResponse(id uuid.UUID, name string, description *string) *PaymentStatusResponse {
	return &PaymentStatusResponse{
		ID:          id,
		Name:        name,
		Description: description,
	}
}

func (r *PaymentStatusRequest) ToDomain() (*domain.PaymentStatus, error) {
	return domain.NewPaymentStatus(r.Name, r.Description)
}
