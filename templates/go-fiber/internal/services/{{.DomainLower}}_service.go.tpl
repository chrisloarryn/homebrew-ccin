package services

import (
	"database/sql"
	"fmt"
	"time"

	"{{.ProjectName}}/internal/models"
)

// {{.DomainTitle}}Service handles business logic for {{.DomainLower}}s
type {{.DomainTitle}}Service struct {
	db *sql.DB
}

// New{{.DomainTitle}}Service creates a new {{.DomainLower}} service
func New{{.DomainTitle}}Service(db *sql.DB) *{{.DomainTitle}}Service {
	return &{{.DomainTitle}}Service{db: db}
}

// GetAll returns all {{.DomainLower}}s
func (s *{{.DomainTitle}}Service) GetAll() ([]models.{{.DomainTitle}}, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM {{.DomainLower}}s ORDER BY created_at DESC`
	
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query {{.DomainLower}}s: %w", err)
	}
	defer rows.Close()

	var {{.DomainLower}}s []models.{{.DomainTitle}}
	for rows.Next() {
		var {{.DomainLower}} models.{{.DomainTitle}}
		err := rows.Scan(&{{.DomainLower}}.ID, &{{.DomainLower}}.Name, &{{.DomainLower}}.Description, &{{.DomainLower}}.CreatedAt, &{{.DomainLower}}.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan {{.DomainLower}}: %w", err)
		}
		{{.DomainLower}}s = append({{.DomainLower}}s, {{.DomainLower}})
	}

	return {{.DomainLower}}s, nil
}

// GetByID returns a {{.DomainLower}} by ID
func (s *{{.DomainTitle}}Service) GetByID(id int) (*models.{{.DomainTitle}}, error) {
	query := `SELECT id, name, description, created_at, updated_at FROM {{.DomainLower}}s WHERE id = $1`
	
	var {{.DomainLower}} models.{{.DomainTitle}}
	err := s.db.QueryRow(query, id).Scan(&{{.DomainLower}}.ID, &{{.DomainLower}}.Name, &{{.DomainLower}}.Description, &{{.DomainLower}}.CreatedAt, &{{.DomainLower}}.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("{{.DomainLower}} not found")
		}
		return nil, fmt.Errorf("failed to get {{.DomainLower}}: %w", err)
	}

	return &{{.DomainLower}}, nil
}

// Create creates a new {{.DomainLower}}
func (s *{{.DomainTitle}}Service) Create(req *models.Create{{.DomainTitle}}Request) (*models.{{.DomainTitle}}, error) {
	query := `INSERT INTO {{.DomainLower}}s (name, description, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4) RETURNING id`
	
	now := time.Now()
	var id int
	err := s.db.QueryRow(query, req.Name, req.Description, now, now).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("failed to create {{.DomainLower}}: %w", err)
	}

	{{.DomainLower}} := &models.{{.DomainTitle}}{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return {{.DomainLower}}, nil
}

// Update updates a {{.DomainLower}}
func (s *{{.DomainTitle}}Service) Update(id int, req *models.Update{{.DomainTitle}}Request) (*models.{{.DomainTitle}}, error) {
	// Check if {{.DomainLower}} exists
	existing, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	if req.Name != "" {
		existing.Name = req.Name
	}
	if req.Description != "" {
		existing.Description = req.Description
	}
	existing.UpdatedAt = time.Now()

	query := `UPDATE {{.DomainLower}}s SET name = $1, description = $2, updated_at = $3 WHERE id = $4`
	_, err = s.db.Exec(query, existing.Name, existing.Description, existing.UpdatedAt, id)
	if err != nil {
		return nil, fmt.Errorf("failed to update {{.DomainLower}}: %w", err)
	}

	return existing, nil
}

// Delete deletes a {{.DomainLower}}
func (s *{{.DomainTitle}}Service) Delete(id int) error {
	// Check if {{.DomainLower}} exists
	_, err := s.GetByID(id)
	if err != nil {
		return err
	}

	query := `DELETE FROM {{.DomainLower}}s WHERE id = $1`
	_, err = s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete {{.DomainLower}}: %w", err)
	}

	return nil
}
