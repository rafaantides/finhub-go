package postgresql

import (
	"context"
	"finhub-go/internal/config"
	"finhub-go/internal/core/domain"
	"finhub-go/internal/core/dto"
	"finhub-go/internal/core/errors"
	"finhub-go/internal/ent"
	"finhub-go/internal/ent/invoice"
	"finhub-go/internal/ent/paymentstatus"
	"finhub-go/internal/utils"
	"finhub-go/internal/utils/pagination"

	"github.com/google/uuid"
)

func (d *PostgreSQL) GetInvoiceByID(ctx context.Context, id uuid.UUID) (*dto.InvoiceResponse, error) {
	row, err := d.Client.Invoice.Query().
		Where(invoice.IDEQ(id)).
		WithStatus().
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return newInvoiceResponse(row)
}

func (d *PostgreSQL) DeleteInvoiceByID(ctx context.Context, id uuid.UUID) error {
	err := d.Client.Invoice.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.ErrNotFound
		}
		return err
	}
	return nil
}

func (d *PostgreSQL) CreateInvoice(ctx context.Context, input domain.Invoice) (*dto.InvoiceResponse, error) {
	created, err := d.Client.Invoice.
		Create().
		SetTitle(input.Title).
		SetAmount(input.Amount).
		SetNillableIssueDate(input.IssueDate).
		SetDueDate(input.DueDate).
		Save(ctx)

	if err != nil {
		return nil, errors.FailedToSave("invoices", err)
	}

	row, err := d.Client.Invoice.
		Query().
		Where(invoice.ID(created.ID)).
		WithStatus().
		Only(ctx)

	if err != nil {
		return nil, errors.FailedToFind("invoice", err)
	}

	return newInvoiceResponse(row)
}

func (d *PostgreSQL) UpdateInvoice(ctx context.Context, input domain.Invoice) (*dto.InvoiceResponse, error) {
	updated, err := d.Client.Invoice.
		UpdateOneID(input.ID).
		SetTitle(input.Title).
		SetAmount(input.Amount).
		SetIssueDate(*input.IssueDate).
		SetDueDate(input.DueDate).
		SetStatusID(*input.StatusID).
		Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrNotFound
		}
		return nil, errors.FailedToSave("invoices", err)
	}

	row, err := d.Client.Invoice.
		Query().
		Where(invoice.ID(updated.ID)).
		WithStatus().
		Only(ctx)

	if err != nil {
		return nil, errors.FailedToFind("invoice", err)
	}

	return newInvoiceResponse(row)
}

func (d *PostgreSQL) ListInvoices(ctx context.Context, flt dto.InvoiceFilters, pgn *pagination.Pagination) ([]dto.InvoiceResponse, error) {
	query := d.Client.Invoice.Query().WithStatus()

	query = applyInvoiceFilters(query, flt, pgn)

	if pgn.OrderDirection == config.OrderAsc {
		query = query.Order(ent.Asc(pgn.OrderBy))
	} else {
		query = query.Order(ent.Desc(pgn.OrderBy))
	}

	query = query.Limit(pgn.PageSize).Offset(pgn.Offset())

	data, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	return newInvoiceResponseList(data)
}

func (d *PostgreSQL) CountInvoices(ctx context.Context, flt dto.InvoiceFilters, pgn *pagination.Pagination) (int, error) {
	query := d.Client.Invoice.Query()
	query = applyInvoiceFilters(query, flt, pgn)

	total, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func mapInvoiceToResponse(row *ent.Invoice) dto.InvoiceResponse {
	response := dto.InvoiceResponse{
		ID:        row.ID,
		Title:     row.Title,
		Amount:    row.Amount,
		IssueDate: utils.SafeToNillableDateTimeString(row.IssueDate),
		DueDate:   utils.ToDateTimeString(row.DueDate),
		CreatedAt: utils.ToDateTimeString(row.CreatedAt),
		UpdatedAt: utils.ToDateTimeString(row.UpdatedAt),
	}

	response.Status = dto.InvoiceStatusResponse{
		ID:   row.Edges.Status.ID,
		Name: row.Edges.Status.Name,
	}

	return response
}

func newInvoiceResponse(row *ent.Invoice) (*dto.InvoiceResponse, error) {
	if row == nil {
		return nil, nil
	}
	response := mapInvoiceToResponse(row)
	return &response, nil
}

func newInvoiceResponseList(rows []*ent.Invoice) ([]dto.InvoiceResponse, error) {
	if rows == nil {
		return nil, nil
	}
	response := make([]dto.InvoiceResponse, 0, len(rows))
	for _, row := range rows {
		response = append(response, mapInvoiceToResponse(row))
	}
	return response, nil
}

func applyInvoiceFilters(query *ent.InvoiceQuery, flt dto.InvoiceFilters, pgn *pagination.Pagination) *ent.InvoiceQuery {
	if pgn.Search != "" {
		query = query.Where(
			invoice.Or(
				invoice.TitleContainsFold(pgn.Search),
				invoice.HasStatusWith(
					paymentstatus.NameContainsFold(pgn.Search),
				),
			),
		)
	}
	if flt.StatusID != nil {
		statusIds := utils.ToUUIDSlice(*flt.StatusID)
		if len(statusIds) > 0 {
			query = query.Where(
				invoice.HasStatusWith(paymentstatus.IDIn(statusIds...)),
			)
		}
	}
	if flt.MinAmount != nil {
		query = query.Where(
			invoice.AmountGTE(*flt.MinAmount),
		)
	}
	if flt.MaxAmount != nil {
		query = query.Where(
			invoice.AmountLTE(*flt.MaxAmount),
		)
	}
	if t := utils.ToDateUnsafe(flt.StartDate); t != nil {
		query = query.Where(invoice.IssueDateGTE(*t))
	}

	if t := utils.ToDateUnsafe(flt.EndDate); t != nil {
		query = query.Where(invoice.IssueDateLTE(*t))
	}

	return query
}
