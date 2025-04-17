package domain

import (
	"finhub-go/internal/core/errors"
	"time"

	"github.com/google/uuid"
)

type PaymentStatus struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewPaymentStatus(
	name string,
	description *string,
) (*PaymentStatus, error) {
	if name == "" {
		return nil, errors.EmptyField("name")
	}
	return &PaymentStatus{
		Name:        name,
		Description: description,
	}, nil
}
