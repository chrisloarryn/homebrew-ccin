package models

import (
	"time"
)

// {{.DomainTitle}} represents a {{.DomainLower}} entity
type {{.DomainTitle}} struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" binding:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Create{{.DomainTitle}}Request represents the request payload for creating a {{.DomainLower}}
type Create{{.DomainTitle}}Request struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// Update{{.DomainTitle}}Request represents the request payload for updating a {{.DomainLower}}
type Update{{.DomainTitle}}Request struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
