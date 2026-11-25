package main

import (
	"context"
	"log"
	"os"

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
	defer categoryDB.Close()
	categoryService := service.NewCategoryService(categoryDB)
	if len(os.Args) == 2 {
		if err = categoryService.AddDefaultCategory(context.Background(), os.Args[1]); err != nil {
			log.Fatalf("failed to add category: %v", err)
		}
	} else if len(os.Args) > 2 {
		log.Fatal("error: should be 1 args: example ./category_service Путешествия")
	} else {
		server := server.NewServer(categoryService, cfg.HttpPort)
		if err := server.Start(); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}
}
