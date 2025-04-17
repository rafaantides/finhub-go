package importer

import (
	"encoding/json"
	"finhub-go/internal/config"
	"finhub-go/internal/core/dto"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func (s *importerService) processDebts(model, action, filename string, rows [][]string, idx map[string]int) error {
	jobID := uuid.New().String()

	baseName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	parts := strings.Split(baseName, "_")
	if len(parts) < 2 {
		return fmt.Errorf("invalid filename format to generate invoice: %s", filename)
	}

	rawTitle := parts[0]
	dueDate := parts[1]

	rawDate := strings.ReplaceAll(parts[1], "-", "")
	invoiceTitle := rawTitle + rawDate

	total := len(rows) - 1

	for i, row := range rows[1:] {
		isFirst := i == 0
		isLast := i == total-1

		var invoice dto.InvoiceRequest

		if isFirst {
			invoice = dto.InvoiceRequest{
				Title:   invoiceTitle,
				DueDate: dueDate,
				Amount:  "0.00",
			}
		}

		debt, err := buildDebtRequest(model, dueDate, row, idx)
		if err != nil {
			return fmt.Errorf("failed to build request: %w", err)
		}

		msg := dto.ImportDebtMessage{
			JobID:        jobID,
			Filename:     filename,
			IsFirstChunk: isFirst,
			IsLastChunk:  isLast,
			Action:       action,
			Data: struct {
				Invoice dto.InvoiceRequest `json:"invoice"`
				Debt    dto.DebtRequest    `json:"debt"`
			}{
				Invoice: invoice,
				Debt:    *debt,
			},
		}

		messageBytes, err := json.Marshal(msg)
		if err != nil {
			return fmt.Errorf("failed to serialize message: %w", err)
		}

		if err := s.mb.SendMessage(config.ResourceDebts, messageBytes); err != nil {
			return err
		}
	}

	return nil
}

func buildDebtRequest(model, dueDate string, row []string, idx map[string]int) (*dto.DebtRequest, error) {
	switch model {
	case config.ModelNubank:
		return nubankToDebtRequest(dueDate, row, idx)
	default:
		return nil, fmt.Errorf("unknown model: %s", model)
	}
}

func nubankToDebtRequest(dueDate string, row []string, idx map[string]int) (*dto.DebtRequest, error) {
	return &dto.DebtRequest{
		DueDate:      &dueDate,
		PurchaseDate: getValue(row, idx, "date"),
		Title:        getValue(row, idx, "title"),
		Amount:       getValue(row, idx, "amount"),
	}, nil
}

func getValue(row []string, idx map[string]int, key string) string {
	if i, ok := idx[key]; ok && i < len(row) {
		return row[i]
	}
	return ""
}
