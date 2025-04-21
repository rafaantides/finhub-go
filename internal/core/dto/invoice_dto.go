package dto

import (
	"finhub-go/internal/core/domain"
	"finhub-go/internal/core/errors"
	"finhub-go/internal/utils"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type InvoiceRequest struct {
	Title     string  `json:"title"`
	IssueDate *string `json:"issue_date"`
	Amount    string  `json:"amount"`
	DueDate   string  `json:"due_date"`
	StatusID  *string `json:"status_id"`
}

type InvoiceFilters struct {
	StatusID  *[]string `form:"status_id"`
	MinAmount *float64  `form:"min_amount"`
	MaxAmount *float64  `form:"max_amount"`
	StartDate *string   `form:"start_date"`
	EndDate   *string   `form:"end_date"`
}

type InvoiceResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	IssueDate *string   `json:"issue_date"`
	DueDate   string    `json:"due_date"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`

	Status InvoiceStatusResponse `json:"status"`
}

type InvoiceStatusResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (r *InvoiceRequest) ToDomain() (*domain.Invoice, error) {
	var statusID *uuid.UUID
	var err error
	if r.StatusID != nil {
		statusID, err = utils.ToNillableUUID(*r.StatusID)
		if err != nil {
			return nil, errors.InvalidParam("status_id", err)
		}
	}

	dueDate, err := utils.ToDateTime(r.DueDate)
	if err != nil {
		return nil, errors.InvalidParam("due_date", err)
	}

	var issueDate *time.Time
	if r.IssueDate != nil {
		issueDate, err = utils.ToNillableDateTime(*r.IssueDate)
		if err != nil {
			return nil, errors.InvalidParam("issue_date", err)
		}
	}

	amount, err := strconv.ParseFloat(r.Amount, 64)
	if err != nil {
		return nil, errors.InvalidParam("amount", err)
	}

	return domain.NewInvoice(statusID, r.Title, amount, dueDate, issueDate)
}
