package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	internalerrors "github.com/OlegLaban/billAI-Category-service/internal/internal_errors"
	"github.com/OlegLaban/billAI-Category-service/internal/models"
	"github.com/OlegLaban/billAI-Category-service/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Handler struct {
	service *service.CategoryService
}

func NewHandler(cs *service.CategoryService) *Handler {
	return &Handler{
		service: cs,
	}
}

func (h *Handler) RegisterRoute() chi.Router {
	r := chi.NewRouter()

	r.Get("/load", h.LoadBaseCategories)

	return r
}

func (h *Handler) LoadBaseCategories(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.Header.Get("X-UserID"))
	if err != nil {
		log.Printf("error to parse id: %v", err)
		writeJson(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid user ID format in header"})
		return
	}
	categories, err := h.service.GetUserCategories(r.Context(), userID)
	if err != nil {
		if errors.Is(err, internalerrors.ErrorUserNotFound) {
			log.Printf("user not found")
			writeJson(w, http.StatusBadRequest, ErrorResponse{Error: "User not found"})
			return
		}
		log.Printf("error in service: %v", err)
		writeJson(w, http.StatusInternalServerError, ErrorResponse{Error: "Server error"})
		return
	}
	type response struct {
		Categories []models.Category `json:"categories"`
	}
	writeJson(w, http.StatusOK, response{Categories: categories})
}

func writeJson(w http.ResponseWriter, status int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
