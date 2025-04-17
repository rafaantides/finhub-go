package postgresql

import (
	"finhub-go/internal/core/domain"
	"finhub-go/internal/core/dto"
	"finhub-go/internal/core/errors"
	"finhub-go/internal/ent"
	"finhub-go/internal/ent/category"
	"finhub-go/internal/ent/debt"
	"finhub-go/internal/ent/invoice"
	"finhub-go/internal/ent/paymentstatus"
	"finhub-go/internal/utils"

	"context"
	"finhub-go/internal/utils/pagination"

	"github.com/google/uuid"
)

func (d *PostgreSQL) GetDebtByID(ctx context.Context, id uuid.UUID) (*dto.DebtResponse, error) {
	row, err := d.Client.Debt.Query().
		Where(debt.IDEQ(id)).
		WithStatus().
		WithCategory().
		WithInvoice().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return newDebtResponse(row)
}

func (d *PostgreSQL) DeleteDebtByID(ctx context.Context, id uuid.UUID) error {
	err := d.Client.Debt.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.ErrNotFound
		}
		return err
	}
	return nil
}

func (d *PostgreSQL) CreateDebt(ctx context.Context, input domain.Debt) (*dto.DebtResponse, error) {
	created, err := d.Client.Debt.
		Create().
		SetTitle(input.Title).
		SetAmount(input.Amount).
		SetNillableDueDate(input.DueDate).
		SetPurchaseDate(input.PurchaseDate).
		SetNillableStatusID(input.StatusID).
		SetNillableInvoiceID(input.InvoiceID).
		SetNillableCategoryID(input.CategoryID).
		Save(ctx)

	if err != nil {
		return nil, errors.FailedToSave("debts", err)
	}

	row, err := d.Client.Debt.Query().
		Where(debt.ID(created.ID)).
		WithStatus().
		WithCategory().
		WithInvoice().
		Only(ctx)

	if err != nil {
		return nil, errors.FailedToFind("debts", err)
	}

	return newDebtResponse(row)
}

func (d *PostgreSQL) UpdateDebt(ctx context.Context, input domain.Debt) (*dto.DebtResponse, error) {
	updated, err := d.Client.Debt.
		UpdateOneID(input.ID).
		SetTitle(input.Title).
		SetAmount(input.Amount).
		SetNillableDueDate(input.DueDate).
		SetPurchaseDate(input.PurchaseDate).
		SetNillableStatusID(input.StatusID).
		SetNillableInvoiceID(input.InvoiceID).
		SetNillableCategoryID(input.CategoryID).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, errors.FailedToSave("debts", err)
	}

	row, err := d.Client.Debt.Query().
		Where(debt.ID(updated.ID)).
		WithStatus().
		WithCategory().
		WithInvoice().
		Only(ctx)

	if err != nil {
		return nil, errors.FailedToFind("debts", err)
	}

	return newDebtResponse(row)
}

func (d *PostgreSQL) ListDebts(ctx context.Context, flt dto.DebtFilters, pgn *pagination.Pagination) ([]dto.DebtResponse, error) {
	query := d.Client.Debt.Query().
		WithStatus().
		WithCategory().
		WithInvoice()

	query = applyDebtFilters(query, flt, pgn)
	query = query.Order(ent.Desc(pgn.OrderBy))
	query = query.Limit(pgn.PageSize).Offset(pgn.Offset())

	data, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	return newDebtResponseList(data)
}

func (d *PostgreSQL) CountDebts(ctx context.Context, flt dto.DebtFilters, pgn *pagination.Pagination) (int, error) {
	query := d.Client.Debt.Query()
	query = applyDebtFilters(query, flt, pgn)

	total, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func mapDebtToResponse(row *ent.Debt) dto.DebtResponse {
	response := dto.DebtResponse{
		ID:           row.ID,
		Title:        row.Title,
		Amount:       row.Amount,
		PurchaseDate: utils.ToDateTimeString(row.PurchaseDate),
		DueDate:      utils.SafeToNillableDateTimeString(row.DueDate),
		CreatedAt:    utils.ToDateTimeString(row.CreatedAt),
		UpdatedAt:    utils.ToDateTimeString(row.UpdatedAt),
	}

	response.Status = dto.DebtStatusResponse{
		ID:   row.Edges.Status.ID,
		Name: row.Edges.Status.Name,
	}

	if row.Edges.Category != nil {
		response.Category = &dto.DebtCategoryResponse{
			ID:   row.Edges.Category.ID,
			Name: row.Edges.Category.Name,
		}
	}

	if row.Edges.Invoice != nil {
		response.Invoice = &dto.DebtInvoiceResponse{
			ID:    row.Edges.Invoice.ID,
			Title: row.Edges.Invoice.Title,
		}
	}

	return response

}

func newDebtResponse(row *ent.Debt) (*dto.DebtResponse, error) {
	if row == nil {
		return nil, nil
	}
	response := mapDebtToResponse(row)
	return &response, nil
}

func newDebtResponseList(rows []*ent.Debt) ([]dto.DebtResponse, error) {
	if rows == nil {
		return nil, nil
	}
	response := make([]dto.DebtResponse, 0, len(rows))
	for _, row := range rows {
		response = append(response, mapDebtToResponse(row))
	}
	return response, nil
}

func applyDebtFilters(query *ent.DebtQuery, flt dto.DebtFilters, pgn *pagination.Pagination) *ent.DebtQuery {
	if pgn.Search != "" {
		query = query.Where(
			debt.Or(
				debt.TitleContainsFold(pgn.Search),
				debt.HasStatusWith(
					paymentstatus.NameContainsFold(pgn.Search),
				),
				debt.HasCategoryWith(
					category.NameContainsFold(pgn.Search),
				),
				debt.HasInvoiceWith(
					invoice.TitleContains(pgn.Search),
				),
			),
		)
	}

	if flt.StatusID != nil {
		statusIds := utils.ToUUIDSlice(*flt.StatusID)
		if len(statusIds) > 0 {
			query = query.Where(
				debt.HasStatusWith(paymentstatus.IDIn(statusIds...)),
			)
		}
	}
	if flt.CategoryID != nil {
		categoryIds := utils.ToUUIDSlice(*flt.CategoryID)
		if len(categoryIds) > 0 {
			query = query.Where(
				debt.HasCategoryWith(category.IDIn(categoryIds...)),
			)
		}
	}
	if flt.InvoiceID != nil {
		invoiceIds := utils.ToUUIDSlice(*flt.InvoiceID)
		if len(invoiceIds) > 0 {
			query = query.Where(
				debt.HasInvoiceWith(invoice.IDIn(invoiceIds...)),
			)
		}
	}
	if flt.MinAmount != nil {
		query = query.Where(
			debt.AmountGTE(*flt.MinAmount),
		)
	}
	if flt.MaxAmount != nil {
		query = query.Where(
			debt.AmountLTE(*flt.MaxAmount),
		)
	}
	if t := utils.ToDateUnsafe(flt.StartDate); t != nil {
		query = query.Where(debt.PurchaseDateGTE(*t))
	}

	if t := utils.ToDateUnsafe(flt.EndDate); t != nil {
		query = query.Where(debt.PurchaseDateLTE(*t))
	}

	return query
}
