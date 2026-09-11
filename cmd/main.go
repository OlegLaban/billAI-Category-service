package main

import (
	"context"
	"log"

	"github.com/OlegLaban/billAI-Category-service/internal/broker"
	"github.com/OlegLaban/billAI-Category-service/internal/config"
	"github.com/OlegLaban/billAI-Category-service/internal/server"
	"github.com/OlegLaban/billAI-Category-service/internal/service"
	"github.com/OlegLaban/billAI-Category-service/internal/storage/postgres"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	categoryDB, err := postgres.NewCategoryDb(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("failed to load DB: %v", err)
	}
	broker, err := broker.NewRedisBroker(cfg.RedisDSN)
	if err != nil {
		log.Fatalf("failed to load redis broker: %v", err)
	}
	categoryService := service.NewCategoryService(categoryDB, broker)
	go func() {
		err = broker.SubscribeUserCreated(context.Background(), categoryService.HandleUserCreated)
		if err != nil {
			log.Fatal(err)
		}
	}()
	server := server.NewServer(categoryService, cfg.HttpPort)
	if err = server.Start(); err != nil {
		log.Fatalf("error in server: %v", err)
	}
}
