package main

import (
	"log"

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
	categoryService := service.NewCategoryService(categoryDB)
	server := server.NewServer(categoryService, cfg.HttpPort)
	if err = server.Start(); err != nil {
		log.Fatalf("error in server: %v", err)
	}
}
