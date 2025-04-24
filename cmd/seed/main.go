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

	// if cfg.SeedPath != "" {
	// 	if err := seedDebts(ctx, postgresRepo, log, cfg.SeedPath); err != nil {
	// 		log.Fatal("Error aseeding debts: %v", err)
	// 	}
	// }

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
		Color       string
	}{
		{
			"Assinaturas e streaming",
			"Gastos recorrentes com serviços digitais",
			"#FF6B6B",
		},
		{
			"Alimentação e delivery",
			"Despesas com delivery, restaurantes, lanches e refeições fora de casa.",
			"#FFA94D",
		},
		{
			"Saúde e bem-estar",
			"Gastos com farmácia, planos de saúde, terapias e autocuidado.",
			"#20C997",
		},
		{
			"Compras",
			"Compras em lojas, marketplaces e produtos online.",
			"#845EF7",
		},
		{
			"Transporte",
			"Custos com apps de corrida e combustíveis.",
			"#339AF0",
		},
		{
			"Vestuário e estética",
			"Despesas com roupas, acessórios, estética e cuidados pessoais.",
			"#FCC419",
		},
		{
			"Mercado e conveniência",
			"Compras em mercados, mercearias e lojas de conveniência.",
			"#69DB7C",
		},
		{
			"Cafés e padarias",
			"Gastos com cafés, padarias e pequenos lanches no dia a dia.",
			"#D97742",
		},
		{
			"Bares e eventos",
			"Despesas com bares, festas, shows, cinema e entretenimento em geral.",
			"#DA77F2",
		},
		{
			"Contas da casa",
			"Gastos essenciais com moradia, como aluguel, condomínio, água, luz e gás.",
			"#FFB3C1",
		},
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
			SetColor(c.Color).
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
