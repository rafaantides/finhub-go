package postgresql

import (
	"context"
	"finhub-go/internal/config"
	"finhub-go/internal/core/domain"
	"finhub-go/internal/core/dto"
	"finhub-go/internal/core/errors"
	"finhub-go/internal/ent"
	"finhub-go/internal/ent/paymentstatus"
	"finhub-go/internal/utils/pagination"

	"github.com/google/uuid"
)

func (d *PostgreSQL) GetPaymentStatusByID(ctx context.Context, id uuid.UUID) (*dto.PaymentStatusResponse, error) {
	row, err := d.Client.PaymentStatus.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return dto.NewPaymentStatusResponse(row.ID, row.Name, row.Description), nil
}

func (d *PostgreSQL) GetPaymentStatusIDByName(ctx context.Context, name *string) (*uuid.UUID, error) {
	if name == nil {
		return nil, nil
	}

	data, err := d.Client.PaymentStatus.Query().Where(paymentstatus.NameEQ(*name)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}

	id := data.ID
	return &id, nil
}

func (d *PostgreSQL) CreatePaymentStatus(ctx context.Context, input domain.PaymentStatus) (*dto.PaymentStatusResponse, error) {
	row, err := d.Client.PaymentStatus.
		Create().
		SetName(input.Name).
		SetNillableDescription(input.Description).
		Save(ctx)

	if err != nil {
		return nil, errors.FailedToSave("payment_status", err)
	}

	return dto.NewPaymentStatusResponse(row.ID, row.Name, row.Description), nil
}

func (d *PostgreSQL) UpdatePaymentStatus(ctx context.Context, id uuid.UUID, input domain.PaymentStatus) (*dto.PaymentStatusResponse, error) {
	row, err := d.Client.PaymentStatus.
		UpdateOneID(id).
		SetName(input.Name).
		SetNillableDescription(input.Description).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, errors.FailedToSave("payment_status", err)
	}

	return dto.NewPaymentStatusResponse(row.ID, row.Name, row.Description), nil
}

func (d *PostgreSQL) DeletePaymentStatusByID(ctx context.Context, id uuid.UUID) error {
	err := d.Client.PaymentStatus.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.ErrNotFound
		}
		return err
	}
	return nil
}

func (d *PostgreSQL) ListPaymentStatus(ctx context.Context, pgn *pagination.Pagination) ([]dto.PaymentStatusResponse, error) {
	query := d.Client.PaymentStatus.Query()

	query = applyPaymentStatusFilters(query, pgn)

	if pgn.OrderDirection == config.OrderAsc {
		query = query.Order(ent.Asc(pgn.OrderBy))
	} else {
		query = query.Order(ent.Desc(pgn.OrderBy))
	}

	query = query.Limit(pgn.PageSize).Offset(pgn.Offset())

	rows, err := query.All(ctx)
	if err != nil {
		return []dto.PaymentStatusResponse{}, err
	}

	response := make([]dto.PaymentStatusResponse, 0, len(rows))
	for _, row := range rows {
		response = append(response, *dto.NewPaymentStatusResponse(row.ID, row.Name, row.Description))
	}
	return response, nil
}

func (d *PostgreSQL) CountPaymentStatus(ctx context.Context, pgn *pagination.Pagination) (int, error) {
	query := d.Client.PaymentStatus.Query()
	query = applyPaymentStatusFilters(query, pgn)

	total, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func applyPaymentStatusFilters(query *ent.PaymentStatusQuery, pgn *pagination.Pagination) *ent.PaymentStatusQuery {
	if pgn.Search != "" {
		query = query.Where(
			paymentstatus.Or(
				paymentstatus.NameContainsFold(pgn.Search),
				paymentstatus.DescriptionContainsFold(pgn.Search),
			),
		)
	}
	return query
}
