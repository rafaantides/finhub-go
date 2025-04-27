package dto

import (
	"finhub-go/internal/core/domain"
	"finhub-go/internal/core/errors"
	"finhub-go/internal/utils"
	"time"

	"github.com/google/uuid"
)

type DebtRequest struct {
	InvoiceID    *string `json:"invoice_id"`
	PurchaseDate string  `json:"purchase_date"`
	DueDate      *string `json:"due_date"`
	Title        string  `json:"title"`
	Amount       float64 `json:"amount"`
	StatusID     *string `json:"status_id"`
	CategoryID   *string `json:"category_id"`
}
type DebtFilters struct {
	// TODO: fazer um bind que funcione com uuid.UUID o ShouldBindQuery n esta reconhecendo o *[]uuid.UUID
	CategoryID *[]string `form:"category_id"`
	StatusID   *[]string `form:"status_id"`
	InvoiceID  *[]string `form:"invoice_id"`
	MinAmount  *float64  `form:"min_amount"`
	MaxAmount  *float64  `form:"max_amount"`
	StartDate  *string   `form:"start_date"`
	EndDate    *string   `form:"end_date"`
}

type DebtResponse struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	Amount       float64   `json:"amount"`
	PurchaseDate string    `json:"purchase_date"`
	DueDate      *string   `json:"due_date"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`

	Status   DebtStatusResponse    `json:"status"`
	Invoice  *DebtInvoiceResponse  `json:"invoice"`
	Category *DebtCategoryResponse `json:"category"`
}

type DebtCategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type DebtStatusResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type DebtInvoiceResponse struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

func (r *DebtRequest) ToDomain() (*domain.Debt, error) {
	purchaseDate, err := utils.ToDateTime(r.PurchaseDate)
	if err != nil {
		return nil, errors.InvalidParam("purchase_date", err)
	}

	var dueDate *time.Time
	if r.DueDate != nil {
		dueDate, err = utils.ToNillableDateTime(*r.DueDate)
		if err != nil {
			return nil, errors.InvalidParam("due_date", err)
		}
	}

	var invoiceID *uuid.UUID
	if r.InvoiceID != nil {
		invoiceID, err = utils.ToNillableUUID(*r.InvoiceID)
		if err != nil {
			return nil, errors.InvalidParam("invoice_id", err)
		}
	}

	var categoryID *uuid.UUID
	if r.CategoryID != nil {
		categoryID, err = utils.ToNillableUUID(*r.CategoryID)
		if err != nil {
			return nil, errors.InvalidParam("category_id", err)
		}
	}

	var statusID *uuid.UUID
	if r.StatusID != nil {
		statusID, err = utils.ToNillableUUID(*r.StatusID)
		if err != nil {
			return nil, errors.InvalidParam("status_id", err)
		}
	}

	return domain.NewDebt(
		invoiceID,
		categoryID,
		statusID,
		r.Title,
		r.Amount,
		purchaseDate,
		dueDate,
	)
}
