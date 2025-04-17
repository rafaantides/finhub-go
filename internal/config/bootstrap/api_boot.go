package bootstrap

import (
	"finhub-go/internal/adapters/messagebus/rabbitmq"
	"finhub-go/internal/adapters/repository/postgresql"
	"finhub-go/internal/config"
	"finhub-go/internal/core/ports/outbound/messagebus"
	"finhub-go/internal/core/ports/outbound/repository"
	"fmt"
)

type APIDeps struct {
	Repo repository.Repository
	Mbus messagebus.MessageBus
}

func InitApi(envPath string) (*APIDeps, error) {

	cfg, err := config.LoadConfig(envPath)
	if err != nil {
		return nil, err
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
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}

	mbus, err := rabbitmq.NewRabbitMQ(
		cfg.MessageBusUser,
		cfg.MessageBusPass,
		cfg.MessageBusHost,
		cfg.MessageBusPort,
	)
	if err != nil {
		repo.Close()
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	return &APIDeps{
		Repo: repo,
		Mbus: mbus,
	}, nil

}
