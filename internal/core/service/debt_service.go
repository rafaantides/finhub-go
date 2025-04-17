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

type debtService struct {
	repo repository.Repository
}

func NewDebtService(repo repository.Repository) inbound.DebtService {
	return &debtService{repo: repo}
}
func (s *debtService) GetDebtByID(ctx context.Context, id uuid.UUID) (*dto.DebtResponse, error) {
	return s.repo.GetDebtByID(ctx, id)
}

func (s *debtService) CreateDebt(ctx context.Context, input domain.Debt) (*dto.DebtResponse, error) {
	return s.repo.CreateDebt(ctx, input)
}

func (s *debtService) UpdateDebt(ctx context.Context, id uuid.UUID, input domain.Debt) (*dto.DebtResponse, error) {
	return s.repo.UpdateDebt(ctx, input)
}

func (s *debtService) DeleteDebtByID(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteDebtByID(ctx, id)
}

func (s *debtService) ListDebts(ctx context.Context, flt dto.DebtFilters, pgn *pagination.Pagination) ([]dto.DebtResponse, int, error) {
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
