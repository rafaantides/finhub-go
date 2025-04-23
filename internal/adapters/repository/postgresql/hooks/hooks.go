package hooks

import (
	"context"
	"fmt"

	"finhub-go/internal/ent"
	"finhub-go/internal/ent/category"
	"finhub-go/internal/ent/paymentstatus"
)

func SetDefaultStatusHook(client *ent.Client) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			status, err := client.PaymentStatus.
				Query().
				Where(paymentstatus.NameEQ("pending")).
				Only(ctx)
			if err != nil {
				return nil, fmt.Errorf("error fetching status 'pending': %w", err)
			}

			switch mut := m.(type) {
			case *ent.DebtMutation:
				if _, exists := mut.StatusID(); !exists {
					mut.SetStatusID(status.ID)
				}
			case *ent.InvoiceMutation:
				if _, exists := mut.StatusID(); !exists {
					mut.SetStatusID(status.ID)
				}
			default:
				return nil, fmt.Errorf("unexpected mutation type: %T", m)
			}

			return next.Mutate(ctx, m)
		})
	}
}

func SetCategoryFromTitleHook(client *ent.Client, categorizer *Categorizer) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			dm, ok := m.(*ent.DebtMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type: %T", m)
			}

			if !dm.Op().Is(ent.OpCreate) {
				return next.Mutate(ctx, m)
			}

			title, exists := dm.Title()
			if !exists {
				return nil, fmt.Errorf("title is required to categorize debt")
			}

			categoryName := categorizer.Categorize(title)

			// Busca a categoria no banco diretamente
			data, err := client.Category.
				Query().
				Where(category.NameEQ(categoryName)).
				Only(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to find category '%s': %w", categoryName, err)
			}

			dm.SetCategoryID(data.ID)

			return next.Mutate(ctx, dm)
		})
	}
}

func UpdateInvoiceAmountHook(client *ent.Client) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			dm, ok := m.(*ent.DebtMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type: %T", m)
			}

			invoiceID, hasInvoice := dm.InvoiceID()
			if !hasInvoice {
				return next.Mutate(ctx, m) // nada a fazer se não tem invoice
			}

			newAmount, hasNewAmount := dm.Amount()
			if !hasNewAmount {
				return next.Mutate(ctx, m) // nada a fazer se não tem amount
			}

			var delta float64

			if dm.Op().Is(ent.OpCreate) {
				delta = newAmount
			} else if dm.Op().Is(ent.OpUpdateOne) {
				// Obtem o valor anterior
				id, _ := dm.ID()
				oldDebt, err := client.Debt.Get(ctx, id)
				if err != nil {
					return nil, fmt.Errorf("failed to get old debt: %w", err)
				}
				delta = newAmount - oldDebt.Amount
			}

			// Atualiza o invoice
			err := client.Invoice.
				UpdateOneID(invoiceID).
				AddAmount(delta).
				Exec(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to update invoice amount: %w", err)
			}

			// continua a mutação
			return next.Mutate(ctx, m)
		})
	}
}
