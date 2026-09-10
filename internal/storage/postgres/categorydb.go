package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/OlegLaban/billAI-Category-service/internal/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type CategoryDB struct {
	Db *sql.DB
}

func NewCategoryDb(dsn string) (*CategoryDB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open data base: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping data base: %w", err)
	}
	return &CategoryDB{Db: db}, nil
}

func (c *CategoryDB) GetUserCategories(ctx context.Context, userID uuid.UUID) (*[]models.Category, error) {
	var categories []models.Category
	query := `SELECT * FROM categories WHERE user_id IS NULL OR user_id = $1`
	rows, err := c.Db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var cat models.Category
		err := rows.Scan(&cat.ID, &cat.UserID, &cat.Name, &cat.IsDefault)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		categories = append(categories, cat)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return &categories, nil
}

func (c *CategoryDB) GetDefaultCategories(ctx context.Context) (*[]models.Category, error) {
	var categories []models.Category
	query := `SELECT * FROM categories WHERE user_id IS NULL`
	rows, err := c.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query categories: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var cat models.Category
		err := rows.Scan(&cat.ID, &cat.UserID, &cat.Name, &cat.IsDefault)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		categories = append(categories, cat)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return &categories, nil
}

func (c *CategoryDB) AddDefaultCategory(ctx context.Context, name string) error {
	query := `INSERT INTO categories (category_name) VALUES ($1)`
	_, err := c.Db.ExecContext(ctx, query, name)
	if err != nil {
		return fmt.Errorf("failed to add category: %w", err)
	}
	return nil
}

func (c *CategoryDB) Close() error {
	return c.Db.Close()
}
