package postgresql

import (
	"context"
	"finhub-go/internal/config"
	"finhub-go/internal/core/domain"
	"finhub-go/internal/core/dto"
	"finhub-go/internal/core/errors"
	"finhub-go/internal/ent"
	"finhub-go/internal/ent/category"
	"finhub-go/internal/utils/pagination"

	"github.com/google/uuid"
)

func (r *PostgreSQL) GetCategoryByID(ctx context.Context, id uuid.UUID) (*dto.CategoryResponse, error) {
	row, err := r.Client.Category.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, errors.FailedToFind("categories", err)
	}
	return dto.NewCategoryResponse(row.ID, row.Name, row.Description), nil
}

func (r *PostgreSQL) GetCategoryIDByName(ctx context.Context, name *string) (*uuid.UUID, error) {
	if name == nil {
		return nil, nil
	}

	data, err := r.Client.Category.Query().Where(category.NameEQ(*name)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, errors.FailedToFind("categories", err)
	}

	id := data.ID
	return &id, nil
}

func (r *PostgreSQL) CreateCategory(ctx context.Context, input domain.Category) (*dto.CategoryResponse, error) {
	row, err := r.Client.Category.
		Create().
		SetName(input.Name).
		SetNillableDescription(input.Description).
		Save(ctx)

	if err != nil {
		return nil, errors.FailedToSave("categories", err)
	}

	return dto.NewCategoryResponse(row.ID, row.Name, row.Description), nil
}

func (r *PostgreSQL) UpdateCategory(ctx context.Context, input domain.Category) (*dto.CategoryResponse, error) {
	row, err := r.Client.Category.
		UpdateOneID(input.ID).
		SetName(input.Name).
		SetNillableDescription(input.Description).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, errors.FailedToUpdate("categories", err)
	}

	return dto.NewCategoryResponse(row.ID, row.Name, row.Description), nil
}

func (r *PostgreSQL) DeleteCategoryByID(ctx context.Context, id uuid.UUID) error {
	err := r.Client.Category.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.ErrNotFound
		}
		return errors.FailedToDelete("categories", err)
	}
	return nil
}

func (r *PostgreSQL) ListCategories(ctx context.Context, pgn *pagination.Pagination) ([]dto.CategoryResponse, error) {
	query := r.Client.Category.Query()
	query = applyCategoryFilters(query, pgn)

	if pgn.OrderDirection == config.OrderAsc {
		query = query.Order(ent.Asc(pgn.OrderBy))
	} else {
		query = query.Order(ent.Desc(pgn.OrderBy))
	}

	query = query.Limit(pgn.PageSize).Offset(pgn.Offset())

	rows, err := query.All(ctx)
	if err != nil {
		return []dto.CategoryResponse{}, err
	}

	response := make([]dto.CategoryResponse, 0, len(rows))
	for _, row := range rows {
		response = append(response, *dto.NewCategoryResponse(row.ID, row.Name, row.Description))
	}
	return response, nil

}

func (r *PostgreSQL) CountCategories(ctx context.Context, pgn *pagination.Pagination) (int, error) {
	query := r.Client.Category.Query()
	query = applyCategoryFilters(query, pgn)

	total, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func applyCategoryFilters(query *ent.CategoryQuery, pgn *pagination.Pagination) *ent.CategoryQuery {
	if pgn.Search != "" {
		query = query.Where(
			category.Or(
				category.NameContainsFold(pgn.Search),
				category.DescriptionContainsFold(pgn.Search),
			),
		)
	}
	return query
}
