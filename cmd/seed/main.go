package main

import (
	"context"
	"encoding/json"
	"finhub-go/internal/adapters/repository/postgresql"
	"finhub-go/internal/config"
	"finhub-go/internal/core/domain"
	"finhub-go/internal/ent/category"
	"finhub-go/internal/ent/paymentstatus"
	"finhub-go/internal/utils/logger"
	"flag"
	"fmt"
	"os"
)

var (
	envPath string
)

func main() {
	flag.StringVar(&envPath, "env", ".env", "Path to .env file")
	flag.Parse()
	startSeed()
}

func startSeed() {
	log := logger.NewLogger("Seed")
	ctx := context.Background()

	log.Start("🌱 Starting database seed... env: %s", envPath)

	cfg, err := config.LoadConfig(envPath)
	if err != nil {
		log.Fatal("%v", err)
	}

	repo, err := postgresql.NewPostgreSQL(
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.SeedPath,
	)

	if err != nil {
		log.Fatal("%v", err)
	}

	postgresRepo, ok := repo.(*postgresql.PostgreSQL)
	if !ok {
		log.Fatal("❌ Type assertion failed: repo is not *postgresql.PostgreSQL")
	}

	if err := seedPaymentStatuses(ctx, postgresRepo, log); err != nil {
		log.Fatal("Error seeding payment statuses: %v", err)
	}

	if err := seedCategories(ctx, postgresRepo, log); err != nil {
		log.Fatal("Error seeding categories: %v", err)
	}

	if cfg.SeedPath != "" {
		if err := seedDebts(ctx, postgresRepo, log, cfg.SeedPath); err != nil {
			log.Fatal("Error aseeding debts: %v", err)
		}
	}

	fmt.Println("✅ Seeding completed successfully!")
}

func seedPaymentStatuses(ctx context.Context, repo *postgresql.PostgreSQL, lg *logger.Logger) error {
	statuses := []struct {
		Name        string
		Description string
	}{
		{"pending", "Pagamento pendente"},
		{"paid", "Pagamento realizado"},
		{"failed", "Pagamento falhou"},
	}

	for _, s := range statuses {
		exists, err := repo.Client.PaymentStatus.Query().Where(paymentstatus.NameEQ(s.Name)).Exist(ctx)

		if err != nil {
			return err
		}
		if exists {
			continue
		}

		_, err = repo.Client.PaymentStatus.
			Create().
			SetName(s.Name).
			SetDescription(s.Description).
			Save(ctx)
		if err != nil {
			return err
		}
		lg.Info("✅ Status created: %s", s.Name)
	}
	return nil
}

func seedCategories(ctx context.Context, repo *postgresql.PostgreSQL, lg *logger.Logger) error {
	categories := []struct {
		Name        string
		Description string
	}{
		{"Assinaturas e Streaming", "Gastos recorrentes com serviços de streaming e assinaturas digitais."},
		{"Alimentação e Delivery", "Despesas com alimentação em restaurantes, delivery e lanches."},
		{"Saúde e Bem-estar", "Compras relacionadas à saúde, farmácia, cosméticos e bem-estar."},
		{"Compras e E-commerce", "Gastos com compras em lojas online e marketplaces."},
		{"Transporte", "Despesas com aplicativos de transporte e mobilidade urbana."},
		{"Vestuario e Estética", "Gastos com roupas, acessórios, estética e cuidados pessoais."},
		{"Mercado e Conveniência", "Compras em supermercados, atacados e lojas de conveniência."},
		{"Cafés e Bares", "Despesas em cafeterias, bares e estabelecimentos similares."},
		{"Eventos e Lazer", "Gastos com cinema, shows, eventos e entretenimento."},
	}

	for _, c := range categories {
		exists, err := repo.Client.Category.Query().Where(category.NameEQ(c.Name)).Exist(ctx)
		if err != nil {
			return err
		}
		if exists {
			continue
		}

		_, err = repo.Client.Category.
			Create().
			SetName(c.Name).
			SetDescription(c.Description).
			Save(ctx)
		if err != nil {
			return err
		}
		lg.Info("✅ Category created: %s", c.Name)
	}
	return nil
}

func seedDebts(ctx context.Context, db *postgresql.PostgreSQL, lg *logger.Logger, seedPath string) error {
	data, err := os.ReadFile(seedPath + "/debts.json")
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo JSON: %w", err)
	}

	var debts []domain.Debt
	if err := json.Unmarshal(data, &debts); err != nil {
		return fmt.Errorf("erro ao parsear JSON: %w", err)
	}

	for _, d := range debts {
		_, err := db.Client.Debt.
			Create().
			SetTitle(d.Title).
			SetAmount(d.Amount).
			SetPurchaseDate(d.PurchaseDate).
			SetDueDate(*d.DueDate).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("erro ao criar dívida '%s': %w", d.Title, err)
		}
		lg.Info("✅ Dívida criada: %s", d.Title)
	}
	return nil
}
