package main

import (
	"flag"
	"fmt"
	"os"

	"finhub-go/internal/config/bootstrap"
	"finhub-go/internal/core/service/consumers"
	"finhub-go/internal/utils/logger"
	"finhub-go/internal/worker"
)

var (
	prefetchCount int
	timeout       int
	envPath       string
)

func main() {
	// Definir flags
	flag.IntVar(&prefetchCount, "prefetch", 10, "Prefetch count for messages")
	flag.IntVar(&timeout, "timeout", 240, "Timeout for message handling in seconds")
	flag.StringVar(&envPath, "env", ".env", "Path to .env file")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Error: you must provide exactly one argument indicating the consumer type")
		os.Exit(1)
	}

	startConsumer(args[0])
}

func startConsumer(resource string) {

	log := logger.NewLogger("Worker")

	boot, err := bootstrap.InitWorker(envPath)
	if err != nil {
		log.Fatal("%v", err)
	}
	defer boot.Repo.Close()
	defer boot.Mbus.Close()
	defer boot.Cache.Close()

	factory, ok := consumers.Registry[resource]
	if !ok {
		log.Fatal("invalid consumer: %s", resource)
	}

	handler := factory(boot)
	w := worker.NewWorker(resource, handler, prefetchCount, timeout, log, boot.Mbus, boot.Cache, boot.Noti)
	w.Start()
}
