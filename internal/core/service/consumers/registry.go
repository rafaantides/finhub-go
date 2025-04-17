package consumers

import (
	"finhub-go/internal/config"
	"finhub-go/internal/config/bootstrap"
	"finhub-go/internal/core/ports/inbound"
	"finhub-go/internal/core/service"
)

type ConsumerFactory func(*bootstrap.WorkerDeps) inbound.MessageProcessor

var Registry = map[string]ConsumerFactory{
	config.ResourceDebts: func(b *bootstrap.WorkerDeps) inbound.MessageProcessor {
		debtService := service.NewDebtService(b.Repo)
		invoiceService := service.NewInvoiceService(b.Repo)
		consumer := NewDebtConsumer(debtService, invoiceService, b.Cache, b.Cfg)
		return consumer.ProcessDebt
	},
}
