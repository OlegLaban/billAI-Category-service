package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OlegLaban/billAI-Category-service/internal/models"
	"github.com/OlegLaban/billAI-Category-service/internal/service"
	"github.com/google/uuid"
)

type Handler struct {
	service *service.CategoryService
}

func NewHandler(cs *service.CategoryService) *Handler {
	return &Handler{
		service: cs,
	}
}

func (h *Handler) RegisterRoute(mux *http.ServeMux) {
	mux.HandleFunc("/load", h.LoadBaseCategories)
}

func (h *Handler) LoadBaseCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error in method, except get", http.StatusMethodNotAllowed)
		return
	}
	userID, err := uuid.Parse(r.Header.Get("X-UserID"))
	if err != nil {
		log.Printf("error to parse id: %v", err)
		http.Error(w, "Invalid user ID format in header", http.StatusBadRequest)
		return
	}
	categories, err := h.service.GetUserCategories(r.Context(), userID)
	if err != nil {
		log.Printf("error in service: %v", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
	}
	type response struct {
		Categories *[]models.Category `json:"categories"`
	}
	fmt.Println(categories)
	resp := response{
		Categories: categories,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
