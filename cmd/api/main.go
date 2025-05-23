package main

import (
	"flag"

	"finhub-go/internal/config/bootstrap"
	"finhub-go/internal/http/routes"
	"finhub-go/internal/utils/logger"
)

var (
	port    string
	envPath string
	debug   bool
)

func main() {
	flag.StringVar(&port, "port", "8080", "Port to run API server on")
	flag.StringVar(&envPath, "env", ".env", "Path to .env file")
	flag.BoolVar(&debug, "debug", false, "Enable debug mode")
	flag.Parse()

	startAPIServer()
}

func startAPIServer() {

	log := logger.NewLogger("API")

	boot, err := bootstrap.InitApi(envPath)
	if err != nil {
		log.Fatal("%v", err)
	}
	defer boot.Repo.Close()
	defer boot.Mbus.Close()

	router := routes.NewRouter(log, boot.Repo, boot.Mbus)
	r := router.Setup(debug)

	log.Start("Starting API server on port %s | env: %s | Debug mode: %v", port, envPath, debug)
	r.Run(":" + port)
}
