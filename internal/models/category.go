package models

import "github.com/google/uuid"

type Category struct {
	Id     int       `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
}
