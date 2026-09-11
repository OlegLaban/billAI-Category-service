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

func (c *CategoryDB) CreateDefaultCategories(ctx context.Context, userID uuid.UUID) error {
	query := `INSERT INTO categories (user_id, category_name) VALUES
	($1, 'Продукты питания'),
    ($1, 'Кафе и рестораны'),
    ($1, 'Для дома и бытовая химия'),
    ($1, 'Здоровье и принадлежности для ухода'),
    ($1, 'Транспорт'),
    ($1, 'Одежда и обувь'),
    ($1, 'Развлечения и хобби'),
    ($1, 'Образование'),
    ($1, 'Домашние животные'),
    ($1, 'Техника и электронника'),
    ($1, 'Жилье и коммунальные услуги'),
    ($1, 'Путешествия')`
	_, err := c.Db.ExecContext(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to create default categories: %w", err)
	}
	return nil
}

func (c *CategoryDB) GetUserCategories(ctx context.Context, userID uuid.UUID) (*[]models.Category, error) {
	query := `SELECT category_name FROM categories
	WHERE user_id = $1`
	rows, err := c.Db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()
	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		err := rows.Scan(&cat.Name)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		categories = append(categories, cat)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iteration failed: %v", err)
	}
	return &categories, nil
}

func (c *CategoryDB) Close() error {
	return c.Db.Close()
}
