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

func (c *CategoryDB) GetUserCategories(ctx context.Context, userID uuid.UUID) ([]models.Category, error) {
	query := `SELECT id, user_id, category_name FROM categories
	WHERE user_id = $1 OR is_default = true`
	rows, err := c.Db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()
	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		err := rows.Scan(&cat.Id, &cat.UserID, &cat.Name, &cat.IsDefault)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		categories = append(categories, cat)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iteration failed: %v", err)
	}
	return categories, nil
}

func (c *CategoryDB) CheckDefaultCategoryName(ctx context.Context, category string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM categories
	WHERE category_name = $1 AND is_default=true)`
	err := c.Db.QueryRowContext(ctx, query, category).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check category: %w", err)
	}
	return exists, nil

}

func (c *CategoryDB) AddDefaultCategory(ctx context.Context, category string) error {
	query := `INSERT INTO categories (category_name, is_default) VALUES($1, $2)`
	_, err := c.Db.ExecContext(ctx, query, category, true)
	if err != nil {
		return fmt.Errorf("failed to add default category: %w", err)
	}
	return nil
}

func (c *CategoryDB) Close() error {
	return c.Db.Close()
}
