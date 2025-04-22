package repository

import (
	"context"
	"finhub-go/internal/core/domain"
	"finhub-go/internal/core/dto"
	"finhub-go/internal/utils/pagination"

	"github.com/google/uuid"
)

type Repository interface {
	// TODO: rever esses docs
	Close()

	GetCategoryByID(ctx context.Context, id uuid.UUID) (*dto.CategoryResponse, error)
	GetCategoryIDByName(ctx context.Context, name *string) (*uuid.UUID, error)
	CreateCategory(ctx context.Context, input domain.Category) (*dto.CategoryResponse, error)
	UpdateCategory(ctx context.Context, input domain.Category) (*dto.CategoryResponse, error)
	DeleteCategoryByID(ctx context.Context, id uuid.UUID) error
	ListCategories(ctx context.Context, pgn *pagination.Pagination) ([]dto.CategoryResponse, error)
	CountCategories(ctx context.Context, pgn *pagination.Pagination) (int, error)

	GetDebtByID(ctx context.Context, id uuid.UUID) (*dto.DebtResponse, error)
	CreateDebt(ctx context.Context, input domain.Debt) (*dto.DebtResponse, error)
	UpdateDebt(ctx context.Context, id uuid.UUID, input domain.Debt) (*dto.DebtResponse, error)
	DeleteDebtByID(ctx context.Context, id uuid.UUID) error
	ListDebts(ctx context.Context, flt dto.DebtFilters, pgn *pagination.Pagination) ([]dto.DebtResponse, error)
	CountDebts(ctx context.Context, flt dto.DebtFilters, pgn *pagination.Pagination) (int, error)
	DebtsSummary(ctx context.Context, flt dto.ChartFilters) ([]dto.SummaryByDate, error)

	GetInvoiceByID(ctx context.Context, id uuid.UUID) (*dto.InvoiceResponse, error)
	CreateInvoice(ctx context.Context, input domain.Invoice) (*dto.InvoiceResponse, error)
	UpdateInvoice(ctx context.Context, input domain.Invoice) (*dto.InvoiceResponse, error)
	DeleteInvoiceByID(ctx context.Context, id uuid.UUID) error
	ListInvoices(ctx context.Context, flt dto.InvoiceFilters, pgn *pagination.Pagination) ([]dto.InvoiceResponse, error)
	CountInvoices(ctx context.Context, flt dto.InvoiceFilters, pgn *pagination.Pagination) (int, error)

	GetPaymentStatusByID(ctx context.Context, id uuid.UUID) (*dto.PaymentStatusResponse, error)
	CreatePaymentStatus(ctx context.Context, input domain.PaymentStatus) (*dto.PaymentStatusResponse, error)
	UpdatePaymentStatus(ctx context.Context, input domain.PaymentStatus) (*dto.PaymentStatusResponse, error)
	DeletePaymentStatusByID(ctx context.Context, id uuid.UUID) error
	ListPaymentStatus(ctx context.Context, pgn *pagination.Pagination) ([]dto.PaymentStatusResponse, error)
	CountPaymentStatus(ctx context.Context, pgn *pagination.Pagination) (int, error)
}
