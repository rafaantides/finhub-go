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

type categoryService struct {
	repo repository.Repository
}

func NewCategoryService(repo repository.Repository) inbound.CategoryService {
	return &categoryService{repo: repo}
}
func (s *categoryService) GetCategoryByID(ctx context.Context, id uuid.UUID) (*dto.CategoryResponse, error) {
	return s.repo.GetCategoryByID(ctx, id)
}

func (s *categoryService) CreateCategory(ctx context.Context, input domain.Category) (*dto.CategoryResponse, error) {
	return s.repo.CreateCategory(ctx, input)
}

func (s *categoryService) UpdateCategory(ctx context.Context, id uuid.UUID, input domain.Category) (*dto.CategoryResponse, error) {
	return s.repo.UpdateCategory(ctx, id, input)
}

func (s *categoryService) DeleteCategoryByID(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteCategoryByID(ctx, id)
}

func (s *categoryService) ListCategories(ctx context.Context, pgn *pagination.Pagination) ([]dto.CategoryResponse, int, error) {
	data, err := s.repo.ListCategories(ctx, pgn)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountCategories(ctx, pgn)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil
}
