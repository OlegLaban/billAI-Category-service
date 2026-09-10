package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/OlegLaban/billAI-Category-service/internal/models"
	"github.com/google/uuid"
)

type CategoryRepository interface {
	GetUserCategories(ctx context.Context, userID uuid.UUID) (*[]models.Category, error)
	AddDefaultCategory(ctx context.Context, name string) error
	GetDefaultCategories(ctx context.Context) (*[]models.Category, error)
}

type CategoryService struct {
	CategoryRepo CategoryRepository
}

func NewCategoryService(cr CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepo: cr}
}

func (cs *CategoryService) GetUserCategories(ctx context.Context, userID uuid.UUID) (*[]models.Category, error) {
	categories, err := cs.CategoryRepo.GetUserCategories(ctx, userID)
	if err != nil || categories == nil {
		return nil, fmt.Errorf("failed to get all categories: %w", err)
	}
	return categories, nil
}

func (cs *CategoryService) AddCategory(ctx context.Context, name string) error {
	name = strings.Join(strings.Fields(name), " ")
	categories, err := cs.CategoryRepo.GetDefaultCategories(ctx)
	if err != nil {
		return fmt.Errorf("error in db to get default categories: %w", err)
	}
	for _, cat := range *categories {
		if strings.EqualFold(cat.Name, name) {
			return fmt.Errorf("category already exist")
		}
	}
	if err := cs.CategoryRepo.AddDefaultCategory(ctx, name); err != nil {
		return fmt.Errorf("failed to add category: %w", err)
	}

	return nil
}
