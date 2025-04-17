package domain

import (
	"finhub-go/internal/core/errors"
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Amount float64   `json:"amount"`
	// Ver a diferença dos tipos das datas e uma forma de colocoar as duas
	IssueDate *time.Time `json:"issue_date"`
	DueDate   time.Time  `json:"due_date"`
	StatusID  *uuid.UUID `json:"status_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewInvoice(
	statusID *uuid.UUID,
	title string,
	amount float64,
	dueDate time.Time,
	issueDate *time.Time,
) (*Invoice, error) {
	if title == "" {
		return nil, errors.EmptyField("name")
	}
	return &Invoice{
		Title:     title,
		Amount:    amount,
		IssueDate: issueDate,
		DueDate:   dueDate,
		StatusID:  statusID,
	}, nil
}
