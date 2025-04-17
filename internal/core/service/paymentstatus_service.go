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

type paymentStatus struct {
	repo repository.Repository
}

func NewPaymentStatusService(repo repository.Repository) inbound.PaymentStatusService {
	return &paymentStatus{repo: repo}
}
func (s *paymentStatus) GetPaymentStatusByID(ctx context.Context, id uuid.UUID) (*dto.PaymentStatusResponse, error) {
	return s.repo.GetPaymentStatusByID(ctx, id)
}

func (s *paymentStatus) CreatePaymentStatus(ctx context.Context, input domain.PaymentStatus) (*dto.PaymentStatusResponse, error) {
	return s.repo.CreatePaymentStatus(ctx, input)
}

func (s *paymentStatus) UpdatePaymentStatus(ctx context.Context, id uuid.UUID, input domain.PaymentStatus) (*dto.PaymentStatusResponse, error) {
	return s.repo.UpdatePaymentStatus(ctx, input)
}

func (s *paymentStatus) DeletePaymentStatusByID(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeletePaymentStatusByID(ctx, id)
}

func (s *paymentStatus) ListPaymentStatus(ctx context.Context, pgn *pagination.Pagination) ([]dto.PaymentStatusResponse, int, error) {
	data, err := s.repo.ListPaymentStatus(ctx, pgn)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountPaymentStatus(ctx, pgn)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}
