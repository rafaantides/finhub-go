package domain

import (
	"finhub-go/internal/core/errors"
	"time"

	"github.com/google/uuid"
)

type Debt struct {
	ID           uuid.UUID  `json:"id"`
	Title        string     `json:"title"`
	Amount       float64    `json:"amount"`
	PurchaseDate time.Time  `json:"purchase_date"`
	DueDate      *time.Time `json:"due_date"`
	InvoiceID    *uuid.UUID `json:"invoice_id"`
	CategoryID   *uuid.UUID `json:"category_id"`
	StatusID     *uuid.UUID `json:"status_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func NewDebt(
	invoiceID, categoryID, statusID *uuid.UUID,
	title string,
	amount float64,
	purchaseDate time.Time,
	dueDate *time.Time,
) (*Debt, error) {
	if title == "" {
		return nil, errors.EmptyField("name")
	}
	return &Debt{
		Title:        title,
		Amount:       amount,
		PurchaseDate: purchaseDate,
		DueDate:      dueDate,
		CategoryID:   categoryID,
		StatusID:     statusID,
		InvoiceID:    invoiceID,
	}, nil
}
