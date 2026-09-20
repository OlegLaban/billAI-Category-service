package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/OlegLaban/billAI-Category-service/internal/handler"
	"github.com/OlegLaban/billAI-Category-service/internal/service"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	server *http.Server
}

func NewServer(cs *service.CategoryService, port string) *Server {
	r := chi.NewRouter()
	categoryHandler := handler.NewHandler(cs)

	r.Mount("/", categoryHandler.RegisterRoute())
	return &Server{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: r,
		},
	}
}

func (s *Server) Start() error {
	log.Printf("Server running on port: %v", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		return fmt.Errorf("server failed: %w", err)
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) Close() error {
	return s.server.Close()
}
