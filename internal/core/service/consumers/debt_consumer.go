package consumers

import (
	"context"
	"encoding/json"
	"finhub-go/internal/config"
	"finhub-go/internal/core/dto"
	"finhub-go/internal/core/ports/inbound"
	"finhub-go/internal/core/ports/outbound/cachestorage"
	"finhub-go/internal/utils/logger"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type DebtConsumer struct {
	debt    inbound.DebtService
	invoice inbound.InvoiceService
	cache   cachestorage.CacheStorage
	cfg     *config.ConfigConsumer
	log     *logger.Logger
}

func NewDebtConsumer(
	debt inbound.DebtService, invoice inbound.InvoiceService,
	cache cachestorage.CacheStorage, cfg *config.ConfigConsumer,
) *DebtConsumer {
	return &DebtConsumer{
		debt:    debt,
		invoice: invoice,
		cache:   cache,
		cfg:     cfg,
		log:     logger.NewLogger("DebtConsumer"),
	}
}

func (d *DebtConsumer) ProcessDebt(messageBody []byte, timeoutSeconds int) (*dto.NotifierMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	var msg dto.ImportDebtMessage
	if err := json.Unmarshal(messageBody, &msg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ImportDebtMessage: %w", err)
	}

	nmsg := &dto.NotifierMessage{
		JobID:       msg.JobID,
		Filename:    msg.Filename,
		IsLastChunk: msg.IsLastChunk,
	}

	d.log.Info("Processing message: %+v", msg)

	switch msg.Action {
	case config.ActionCreate:
		debtInput, err := msg.Data.Debt.ToDomain()
		if err != nil {
			return nmsg, fmt.Errorf("failed to parse debt: %w", err)
		}

		if d.shouldSkipTitle(debtInput.Title) {
			d.log.Info("Skipping title: %s", debtInput.Title)
			return nmsg, nil
		}

		jobKey := fmt.Sprintf("invoice:%s", msg.JobID)
		var invoiceID uuid.UUID

		if msg.IsFirstChunk {
			ok, err := d.cache.SetNX(
				ctx,
				jobKey,
				"creating",
				time.Duration(d.cfg.InvoiceCacheTTLMin)*time.Minute,
			)
			if err != nil {
				return nmsg, fmt.Errorf("failed to set cache (SetNX): %w", err)
			}

			if ok {
				invoiceInput, err := msg.Data.Invoice.ToDomain()
				if err != nil {
					return nmsg, fmt.Errorf("failed to parse invoice: %w", err)
				}

				dataInvoice, err := d.invoice.CreateInvoice(ctx, *invoiceInput)
				if err != nil {
					return nmsg, fmt.Errorf("failed to create invoice: %w", err)
				}
				invoiceID = dataInvoice.ID

				if _, err = d.cache.Set(
					ctx,
					jobKey,
					invoiceID.String(),
					time.Duration(d.cfg.InvoiceCacheTTLMin)*time.Minute,
				); err != nil {
					return nmsg, fmt.Errorf("failed to cache invoice ID: %w", err)
				}
			} else {
				invoiceID, err = d.waitForInvoiceID(ctx, jobKey)
				if err != nil {
					return nmsg, fmt.Errorf("waiting for invoice creation: %w", err)
				}
			}
		} else {
			invoiceID, err = d.waitForInvoiceID(ctx, jobKey)
			if err != nil {
				return nmsg, fmt.Errorf("waiting for invoice_id: %w", err)
			}
		}

		debtInput.InvoiceID = &invoiceID

		if _, err := d.debt.CreateDebt(ctx, *debtInput); err != nil {
			return nmsg, fmt.Errorf("failed to create debt: %w", err)
		}

	default:
		return nmsg, fmt.Errorf("invalid action: %s", msg.Action)
	}

	return nmsg, nil
}

func (d *DebtConsumer) waitForInvoiceID(ctx context.Context, key string) (uuid.UUID, error) {
	val, err := d.cache.WaitForCacheValue(
		ctx,
		key,
		time.Duration(d.cfg.PollIntervalMs)*time.Millisecond,
		time.Duration(d.cfg.WaitForInvoiceLimit)*time.Second,
		func(val string) (bool, error) {
			return val != "creating", nil
		},
	)

	if err != nil {
		return uuid.Nil, err
	}

	id, parseErr := uuid.Parse(val)
	if parseErr != nil {
		// TODO: rever erros de parse
		return uuid.Nil, fmt.Errorf("error parsing UUID: %w", parseErr)
	}

	return id, nil
}

func (d *DebtConsumer) shouldSkipTitle(title string) bool {
	for _, skip := range d.cfg.SkipTitles {
		if strings.EqualFold(skip, title) {
			return true
		}
	}
	return false
}
