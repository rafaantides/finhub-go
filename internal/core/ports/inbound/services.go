package inbound

import (
	"context"
	"finhub-go/internal/core/domain"
	"finhub-go/internal/core/dto"
	"finhub-go/internal/utils/pagination"
	"mime/multipart"

	"github.com/google/uuid"
)

type CategoryService interface {
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*dto.CategoryResponse, error)
	CreateCategory(ctx context.Context, input domain.Category) (*dto.CategoryResponse, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, input domain.Category) (*dto.CategoryResponse, error)
	DeleteCategoryByID(ctx context.Context, id uuid.UUID) error
	ListCategories(ctx context.Context, pgn *pagination.Pagination) ([]dto.CategoryResponse, int, error)
}

type DebtService interface {
	GetDebtByID(ctx context.Context, id uuid.UUID) (*dto.DebtResponse, error)
	CreateDebt(ctx context.Context, input domain.Debt) (*dto.DebtResponse, error)
	UpdateDebt(ctx context.Context, id uuid.UUID, input domain.Debt) (*dto.DebtResponse, error)
	DeleteDebtByID(ctx context.Context, id uuid.UUID) error
	ListDebts(ctx context.Context, flt dto.DebtFilters, pgn *pagination.Pagination) ([]dto.DebtResponse, int, error)
}

type InvoiceService interface {
	GetInvoiceByID(ctx context.Context, id uuid.UUID) (*dto.InvoiceResponse, error)
	CreateInvoice(ctx context.Context, input domain.Invoice) (*dto.InvoiceResponse, error)
	UpdateInvoice(ctx context.Context, id uuid.UUID, input domain.Invoice) (*dto.InvoiceResponse, error)
	DeleteInvoiceByID(ctx context.Context, id uuid.UUID) error
	ListInvoices(ctx context.Context, flt dto.InvoiceFilters, pgn *pagination.Pagination) ([]dto.InvoiceResponse, int, error)
}

type PaymentStatusService interface {
	GetPaymentStatusByID(ctx context.Context, id uuid.UUID) (*dto.PaymentStatusResponse, error)
	CreatePaymentStatus(ctx context.Context, input domain.PaymentStatus) (*dto.PaymentStatusResponse, error)
	UpdatePaymentStatus(ctx context.Context, id uuid.UUID, input domain.PaymentStatus) (*dto.PaymentStatusResponse, error)
	DeletePaymentStatusByID(ctx context.Context, id uuid.UUID) error
	ListPaymentStatus(ctx context.Context, pgn *pagination.Pagination) ([]dto.PaymentStatusResponse, int, error)
}

type ImporterService interface {
	ImportFile(resource, model, action string, file multipart.File, fileHeader *multipart.FileHeader) error
}
