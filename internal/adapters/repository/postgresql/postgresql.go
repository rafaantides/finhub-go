package postgresql

import (
	"finhub-go/internal/adapters/repository/postgresql/hooks"
	"finhub-go/internal/core/ports/outbound/repository"
	"finhub-go/internal/ent"
	"finhub-go/internal/utils/logger"
	"fmt"
	"net/url"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	log    *logger.Logger
	Client *ent.Client
}

func NewPostgreSQL(user, password, host, port, database, SeedPath string) (repository.Repository, error) {
	log := logger.NewLogger("PostgreSQL")

	dbURI := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		url.QueryEscape(user),
		url.QueryEscape(password),
		host,
		port,
		database,
	)

	drv, err := sql.Open(dialect.Postgres, dbURI)
	if err != nil {
		return nil, err
	}

	sqlDB := drv.DB()
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	client := ent.NewClient(ent.Driver(drv))
	categorizer, err := hooks.NewCategorizer(SeedPath)

	if err != nil {
		return nil, err
	}

	client.Debt.Use(
		hooks.SetDefaultStatusHook(client),
		hooks.UpdateInvoiceAmountHook(client),
		hooks.SetCategoryFromTitleHook(client, categorizer),
	)

	client.Invoice.Use(
		hooks.SetDefaultStatusHook(client),
	)

	log.Start("Host: %s:%s | User: %s | DB: %s", host, port, user, database)

	return &PostgreSQL{Client: client, log: log}, nil
}

func (d *PostgreSQL) Close() {
	if err := d.Client.Close(); err != nil {
		d.log.Error("%v", err)
	} else {
		d.log.Info("Database connection closed.")
	}
}
