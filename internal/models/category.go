package models

import "github.com/google/uuid"

type Category struct {
	ID        int       `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	IsDefault bool      `json:"is_default"`
	Name      string    `json:"name"`
}
