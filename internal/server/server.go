package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/OlegLaban/billAI-Category-service/internal/handler"
	"github.com/OlegLaban/billAI-Category-service/internal/service"
)

type Server struct {
	server *http.Server
}

func NewServer(cs *service.CategoryService, port string) *Server {
	mux := http.NewServeMux()
	handler := handler.NewHandler(cs)
	handler.RegisterRoute(mux)
	return &Server{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: mux,
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

func (s *Server) Close() error {
	return s.server.Close()
}
