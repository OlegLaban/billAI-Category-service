package models

import "github.com/google/uuid"

type Category struct {
	Id     int
	UserID uuid.UUID
	Name   string `json:"name"`
}
