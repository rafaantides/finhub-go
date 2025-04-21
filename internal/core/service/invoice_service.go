package service

import (
	"context"

	"finhub-go/internal/core/domain"
	"finhub-go/internal/core/dto"
	"finhub-go/internal/core/ports/inbound"
	"finhub-go/internal/core/ports/outbound/repository"
	"finhub-go/internal/utils/pagination"

	"github.com/google/uuid"
)

type invoiceService struct {
	repo repository.Repository
}

func NewInvoiceService(repo repository.Repository) inbound.InvoiceService {
	return &invoiceService{repo: repo}
}
func (s *invoiceService) GetInvoiceByID(ctx context.Context, id uuid.UUID) (*dto.InvoiceResponse, error) {
	return s.repo.GetInvoiceByID(ctx, id)
}

func (s *invoiceService) CreateInvoice(ctx context.Context, input domain.Invoice) (*dto.InvoiceResponse, error) {
	return s.repo.CreateInvoice(ctx, input)
}

func (s *invoiceService) UpdateInvoice(ctx context.Context, id uuid.UUID, input domain.Invoice) (*dto.InvoiceResponse, error) {
	return s.repo.UpdateInvoice(ctx, input)
}

func (s *invoiceService) DeleteInvoiceByID(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteInvoiceByID(ctx, id)
}

func (s *invoiceService) ListInvoices(ctx context.Context, flt dto.InvoiceFilters, pgn *pagination.Pagination) ([]dto.InvoiceResponse, int, error) {
	data, err := s.repo.ListInvoices(ctx, flt, pgn)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountInvoices(ctx, flt, pgn)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (s *invoiceService) ListInvoiceDebts(ctx context.Context, id uuid.UUID, flt dto.DebtFilters, pgn *pagination.Pagination) ([]dto.DebtResponse, int, error) {

	flt.InvoiceID = &[]string{id.String()}

	data, err := s.repo.ListDebts(ctx, flt, pgn)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountDebts(ctx, flt, pgn)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil

}
