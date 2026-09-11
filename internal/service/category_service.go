package service

import (
	"context"
	"fmt"

	"github.com/OlegLaban/billAI-Category-service/internal/models"
	"github.com/google/uuid"
)

type CategoryRepository interface {
	CreateDefaultCategories(ctx context.Context, userID uuid.UUID) error
	GetUserCategories(ctx context.Context, userID uuid.UUID) (*[]models.Category, error)
}

type CategoryBroker interface {
	SubscribeUserCreated(ctx context.Context, handler func(context.Context, uuid.UUID) error) error
}

type CategoryService struct {
	CategoryRepo CategoryRepository
	Broker       CategoryBroker
}

func NewCategoryService(cr CategoryRepository, cb CategoryBroker) *CategoryService {
	return &CategoryService{CategoryRepo: cr, Broker: cb}
}

func (cs *CategoryService) GetUserCategories(ctx context.Context, userID uuid.UUID) (*[]models.Category, error) {
	categories, err := cs.CategoryRepo.GetUserCategories(ctx, userID)
	if err != nil || categories == nil {
		return nil, fmt.Errorf("failed to get all categories: %w", err)
	}
	return categories, nil
}

func (cs *CategoryService) HandleUserCreated(ctx context.Context, userID uuid.UUID) error {
	return cs.CategoryRepo.CreateDefaultCategories(ctx, userID)
}
