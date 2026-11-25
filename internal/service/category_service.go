package service

import (
	"context"
	"fmt"
	"strings"
	"unicode"

	internalerrors "github.com/OlegLaban/billAI-Category-service/internal/internal_errors"
	"github.com/OlegLaban/billAI-Category-service/internal/models"
	"github.com/google/uuid"
)

type CategoryRepository interface {
	GetUserCategories(ctx context.Context, userID uuid.UUID) ([]models.Category, error)
	AddDefaultCategory(ctx context.Context, category string) error
	CheckDefaultCategoryName(ctx context.Context, category string) (bool, error)
}

type CategoryService struct {
	CategoryRepo CategoryRepository
}

func NewCategoryService(cr CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepo: cr}
}

func (cs *CategoryService) GetUserCategories(ctx context.Context, userID uuid.UUID) ([]models.Category, error) {
	categories, err := cs.CategoryRepo.GetUserCategories(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get all categories: %w", err)
	}
	if categories == nil {
		return nil, internalerrors.ErrorUserNotFound
	}
	return categories, nil
}

func (cs *CategoryService) AddDefaultCategory(ctx context.Context, category string) error {
	arr := strings.Fields(category)
	category = strings.Join(arr, " ")
	runes := []rune(category)
	runes[0] = unicode.ToUpper(runes[0])
	category = string(runes)
	exists, err := cs.CategoryRepo.CheckDefaultCategoryName(ctx, category)
	if err != nil {
		return fmt.Errorf("failed in categorydb: %w", err)
	}
	if !exists {
		if err := cs.CategoryRepo.AddDefaultCategory(ctx, category); err != nil {
			return fmt.Errorf("failed in categorydb: %w", err)
		}
		return nil
	}
	return internalerrors.ErrorCategoryExist
}
