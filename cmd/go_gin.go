package cmd

import (
	"fmt"
	"path/filepath"
)

func generateGoGinProject(projectName, domainName, gcpProject string, grpc bool) error {
	if err := createProjectDir(projectName); err != nil {
		return err
	}

	// Set default domain if not provided
	if domainName == "" {
		domainName = "item"
	}

	// Generate go.mod
	goMod := replaceTemplateVars(`module {{PROJECT_NAME}}

go 1.24.4

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/swag v1.16.2
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
	gorm.io/gorm v1.25.12
	gorm.io/driver/postgres v1.5.9
	cloud.google.com/go/monitoring v1.20.4
	cloud.google.com/go/logging v1.11.0
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
)`, projectName, domainName, gcpProject)

	if err := createFile(filepath.Join(projectName, "go.mod"), goMod); err != nil {
		return err
	}

	// Generate main.go
	mainGo := generateMainGoFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, "cmd/server/main.go"), mainGo); err != nil {
		return err
	}

	// Generate entity/model
	model := generateModelFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, fmt.Sprintf("internal/models/%s.go", domainName)), model); err != nil {
		return err
	}

	// Generate repository
	repository := generateRepositoryFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, fmt.Sprintf("internal/repository/%s_repository.go", domainName)), repository); err != nil {
		return err
	}

	// Generate service
	service := generateServiceFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, fmt.Sprintf("internal/service/%s_service.go", domainName)), service); err != nil {
		return err
	}

	// Generate handler
	handler := generateHandlerFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, fmt.Sprintf("internal/handlers/%s_handler.go", domainName)), handler); err != nil {
		return err
	}

	// Generate middleware
	middleware := generateMiddlewareFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, "internal/middleware/metrics.go"), middleware); err != nil {
		return err
	}

	// Generate config
	config := generateConfigFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, "internal/config/config.go"), config); err != nil {
		return err
	}

	// Generate router
	router := generateRouterFile(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, "internal/router/router.go"), router); err != nil {
		return err
	}

	// Generate gRPC files if requested
	if grpc {
		if err := generateGRPCFiles(projectName, domainName, gcpProject); err != nil {
			return err
		}
	}

	// Generate Dockerfile
	dockerfile := generateDockerfileGo(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, "Dockerfile"), dockerfile); err != nil {
		return err
	}

	// Generate Makefile
	makefile := generateMakefileGo(projectName, domainName, gcpProject, grpc)
	if err := createFile(filepath.Join(projectName, "Makefile"), makefile); err != nil {
		return err
	}

	// Generate .env.example
	envExample := generateEnvExampleGo(projectName, domainName, gcpProject)
	if err := createFile(filepath.Join(projectName, ".env.example"), envExample); err != nil {
		return err
	}

	// Generate README.md
	readme := generateReadmeGo(projectName, domainName, gcpProject, grpc)
	if err := createFile(filepath.Join(projectName, "README.md"), readme); err != nil {
		return err
	}

	return nil
}

func generateMainGoFile(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`package main

import (
	"log"
	"{{PROJECT_NAME}}/internal/config"
	"{{PROJECT_NAME}}/internal/router"
	"{{PROJECT_NAME}}/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title {{PROJECT_NAME}} API
// @version 1.0
// @description {{PROJECT_NAME}} CRUD API with Gin framework
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg := config.Load()

	// Setup Gin mode
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	r := gin.New()

	// Add global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.GCPMetrics(cfg.GCPProjectID))

	// Setup routes
	router.SetupRoutes(r, cfg)

	// Start server
	log.Printf("Starting server on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}`, projectName, domainName, gcpProject)
}

func generateModelFile(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// {{DOMAIN_TITLE}} represents the {{DOMAIN_LOWER}} entity
type {{DOMAIN_TITLE}} struct {
	ID          uuid.UUID      'json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"'
	Name        string         'json:"name" gorm:"not null" validate:"required"'
	Description *string        'json:"description,omitempty"'
	IsActive    bool           'json:"is_active" gorm:"default:true"'
	CreatedAt   time.Time      'json:"created_at"'
	UpdatedAt   time.Time      'json:"updated_at"'
	DeletedAt   gorm.DeletedAt 'json:"-" gorm:"index"'
}

// Create{{DOMAIN_TITLE}}Request represents the request payload for creating a {{DOMAIN_LOWER}}
type Create{{DOMAIN_TITLE}}Request struct {
	Name        string  'json:"name" validate:"required" example:"Sample {{DOMAIN_TITLE}}"'
	Description *string 'json:"description,omitempty" example:"Description of the {{DOMAIN_LOWER}}"'
	IsActive    *bool   'json:"is_active,omitempty" example:"true"'
}

// Update{{DOMAIN_TITLE}}Request represents the request payload for updating a {{DOMAIN_LOWER}}
type Update{{DOMAIN_TITLE}}Request struct {
	Name        *string 'json:"name,omitempty" example:"Updated {{DOMAIN_TITLE}}"'
	Description *string 'json:"description,omitempty" example:"Updated description"'
	IsActive    *bool   'json:"is_active,omitempty" example:"false"'
}

// {{DOMAIN_TITLE}}Response represents the response payload for {{DOMAIN_LOWER}} operations
type {{DOMAIN_TITLE}}Response struct {
	ID          uuid.UUID 'json:"id"'
	Name        string    'json:"name"'
	Description *string   'json:"description,omitempty"'
	IsActive    bool      'json:"is_active"'
	CreatedAt   time.Time 'json:"created_at"'
	UpdatedAt   time.Time 'json:"updated_at"'
}

// TableName returns the table name for the {{DOMAIN_TITLE}} model
func ({{DOMAIN_TITLE}}) TableName() string {
	return "{{DOMAIN_LOWER}}s"
}`, projectName, domainName, gcpProject)
}

func generateRepositoryFile(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`package repository

import (
	"{{PROJECT_NAME}}/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// {{DOMAIN_TITLE}}Repository interface defines the contract for {{DOMAIN_LOWER}} data operations
type {{DOMAIN_TITLE}}Repository interface {
	Create({{DOMAIN_LOWER}} *models.{{DOMAIN_TITLE}}) error
	FindAll() ([]models.{{DOMAIN_TITLE}}, error)
	FindByID(id uuid.UUID) (*models.{{DOMAIN_TITLE}}, error)
	Update({{DOMAIN_LOWER}} *models.{{DOMAIN_TITLE}}) error
	Delete(id uuid.UUID) error
}

// {{DOMAIN_LOWER}}Repository implements {{DOMAIN_TITLE}}Repository interface
type {{DOMAIN_LOWER}}Repository struct {
	db *gorm.DB
}

// New{{DOMAIN_TITLE}}Repository creates a new instance of {{DOMAIN_LOWER}}Repository
func New{{DOMAIN_TITLE}}Repository(db *gorm.DB) {{DOMAIN_TITLE}}Repository {
	return &{{DOMAIN_LOWER}}Repository{db: db}
}

// Create creates a new {{DOMAIN_LOWER}} in the database
func (r *{{DOMAIN_LOWER}}Repository) Create({{DOMAIN_LOWER}} *models.{{DOMAIN_TITLE}}) error {
	return r.db.Create({{DOMAIN_LOWER}}).Error
}

// FindAll retrieves all active {{DOMAIN_LOWER}}s from the database
func (r *{{DOMAIN_LOWER}}Repository) FindAll() ([]models.{{DOMAIN_TITLE}}, error) {
	var {{DOMAIN_LOWER}}s []models.{{DOMAIN_TITLE}}
	err := r.db.Where("is_active = ?", true).Find(&{{DOMAIN_LOWER}}s).Error
	return {{DOMAIN_LOWER}}s, err
}

// FindByID retrieves a {{DOMAIN_LOWER}} by its ID
func (r *{{DOMAIN_LOWER}}Repository) FindByID(id uuid.UUID) (*models.{{DOMAIN_TITLE}}, error) {
	var {{DOMAIN_LOWER}} models.{{DOMAIN_TITLE}}
	err := r.db.Where("id = ? AND is_active = ?", id, true).First(&{{DOMAIN_LOWER}}).Error
	if err != nil {
		return nil, err
	}
	return &{{DOMAIN_LOWER}}, nil
}

// Update updates an existing {{DOMAIN_LOWER}} in the database
func (r *{{DOMAIN_LOWER}}Repository) Update({{DOMAIN_LOWER}} *models.{{DOMAIN_TITLE}}) error {
	return r.db.Save({{DOMAIN_LOWER}}).Error
}

// Delete soft deletes a {{DOMAIN_LOWER}} by setting is_active to false
func (r *{{DOMAIN_LOWER}}Repository) Delete(id uuid.UUID) error {
	return r.db.Model(&models.{{DOMAIN_TITLE}}{}).Where("id = ?", id).Update("is_active", false).Error
}`, projectName, domainName, gcpProject)
}

func generateServiceFile(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`package service

import (
	"errors"
	"{{PROJECT_NAME}}/internal/models"
	"{{PROJECT_NAME}}/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// {{DOMAIN_TITLE}}Service interface defines the business logic contract for {{DOMAIN_LOWER}} operations
type {{DOMAIN_TITLE}}Service interface {
	Create(req *models.Create{{DOMAIN_TITLE}}Request) (*models.{{DOMAIN_TITLE}}Response, error)
	GetAll() ([]models.{{DOMAIN_TITLE}}Response, error)
	GetByID(id uuid.UUID) (*models.{{DOMAIN_TITLE}}Response, error)
	Update(id uuid.UUID, req *models.Update{{DOMAIN_TITLE}}Request) (*models.{{DOMAIN_TITLE}}Response, error)
	Delete(id uuid.UUID) error
}

// {{DOMAIN_LOWER}}Service implements {{DOMAIN_TITLE}}Service interface
type {{DOMAIN_LOWER}}Service struct {
	repo repository.{{DOMAIN_TITLE}}Repository
}

// New{{DOMAIN_TITLE}}Service creates a new instance of {{DOMAIN_LOWER}}Service
func New{{DOMAIN_TITLE}}Service(repo repository.{{DOMAIN_TITLE}}Repository) {{DOMAIN_TITLE}}Service {
	return &{{DOMAIN_LOWER}}Service{repo: repo}
}

// Create creates a new {{DOMAIN_LOWER}}
func (s *{{DOMAIN_LOWER}}Service) Create(req *models.Create{{DOMAIN_TITLE}}Request) (*models.{{DOMAIN_TITLE}}Response, error) {
	{{DOMAIN_LOWER}} := &models.{{DOMAIN_TITLE}}{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    true,
	}

	if req.IsActive != nil {
		{{DOMAIN_LOWER}}.IsActive = *req.IsActive
	}

	if err := s.repo.Create({{DOMAIN_LOWER}}); err != nil {
		return nil, err
	}

	return s.modelToResponse({{DOMAIN_LOWER}}), nil
}

// GetAll retrieves all {{DOMAIN_LOWER}}s
func (s *{{DOMAIN_LOWER}}Service) GetAll() ([]models.{{DOMAIN_TITLE}}Response, error) {
	{{DOMAIN_LOWER}}s, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]models.{{DOMAIN_TITLE}}Response, len({{DOMAIN_LOWER}}s))
	for i, {{DOMAIN_LOWER}} := range {{DOMAIN_LOWER}}s {
		responses[i] = *s.modelToResponse(&{{DOMAIN_LOWER}})
	}

	return responses, nil
}

// GetByID retrieves a {{DOMAIN_LOWER}} by its ID
func (s *{{DOMAIN_LOWER}}Service) GetByID(id uuid.UUID) (*models.{{DOMAIN_TITLE}}Response, error) {
	{{DOMAIN_LOWER}}, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("{{DOMAIN_LOWER}} not found")
		}
		return nil, err
	}

	return s.modelToResponse({{DOMAIN_LOWER}}), nil
}

// Update updates an existing {{DOMAIN_LOWER}}
func (s *{{DOMAIN_LOWER}}Service) Update(id uuid.UUID, req *models.Update{{DOMAIN_TITLE}}Request) (*models.{{DOMAIN_TITLE}}Response, error) {
	{{DOMAIN_LOWER}}, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("{{DOMAIN_LOWER}} not found")
		}
		return nil, err
	}

	// Update fields if provided
	if req.Name != nil {
		{{DOMAIN_LOWER}}.Name = *req.Name
	}
	if req.Description != nil {
		{{DOMAIN_LOWER}}.Description = req.Description
	}
	if req.IsActive != nil {
		{{DOMAIN_LOWER}}.IsActive = *req.IsActive
	}

	if err := s.repo.Update({{DOMAIN_LOWER}}); err != nil {
		return nil, err
	}

	return s.modelToResponse({{DOMAIN_LOWER}}), nil
}

// Delete deletes a {{DOMAIN_LOWER}}
func (s *{{DOMAIN_LOWER}}Service) Delete(id uuid.UUID) error {
	// Check if {{DOMAIN_LOWER}} exists
	_, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("{{DOMAIN_LOWER}} not found")
		}
		return err
	}

	return s.repo.Delete(id)
}

// modelToResponse converts a {{DOMAIN_TITLE}} model to response format
func (s *{{DOMAIN_LOWER}}Service) modelToResponse({{DOMAIN_LOWER}} *models.{{DOMAIN_TITLE}}) *models.{{DOMAIN_TITLE}}Response {
	return &models.{{DOMAIN_TITLE}}Response{
		ID:          {{DOMAIN_LOWER}}.ID,
		Name:        {{DOMAIN_LOWER}}.Name,
		Description: {{DOMAIN_LOWER}}.Description,
		IsActive:    {{DOMAIN_LOWER}}.IsActive,
		CreatedAt:   {{DOMAIN_LOWER}}.CreatedAt,
		UpdatedAt:   {{DOMAIN_LOWER}}.UpdatedAt,
	}
}`, projectName, domainName, gcpProject)
}

func generateHandlerFile(projectName, domainName, gcpProject string) string {
	return replaceTemplateVars(`package handlers

import (
	"net/http"
	"{{PROJECT_NAME}}/internal/models"
	"{{PROJECT_NAME}}/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// {{DOMAIN_TITLE}}Handler handles HTTP requests for {{DOMAIN_LOWER}} operations
type {{DOMAIN_TITLE}}Handler struct {
	service service.{{DOMAIN_TITLE}}Service
}

// New{{DOMAIN_TITLE}}Handler creates a new instance of {{DOMAIN_TITLE}}Handler
func New{{DOMAIN_TITLE}}Handler(service service.{{DOMAIN_TITLE}}Service) *{{DOMAIN_TITLE}}Handler {
	return &{{DOMAIN_TITLE}}Handler{service: service}
}

// Create{{DOMAIN_TITLE}} godoc
// @Summary Create a new {{DOMAIN_LOWER}}
// @Description Create a new {{DOMAIN_LOWER}} with the input payload
// @Tags {{DOMAIN_LOWER}}s
// @Accept json
// @Produce json
// @Param {{DOMAIN_LOWER}} body models.Create{{DOMAIN_TITLE}}Request true "{{DOMAIN_TITLE}} data"
// @Success 201 {object} models.{{DOMAIN_TITLE}}Response
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /{{DOMAIN_LOWER}}s [post]
func (h *{{DOMAIN_TITLE}}Handler) Create{{DOMAIN_TITLE}}(c *gin.Context) {
	var req models.Create{{DOMAIN_TITLE}}Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	{{DOMAIN_LOWER}}, err := h.service.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, {{DOMAIN_LOWER}})
}

// Get{{DOMAIN_TITLE}}s godoc
// @Summary Get all {{DOMAIN_LOWER}}s
// @Description Get a list of all {{DOMAIN_LOWER}}s
// @Tags {{DOMAIN_LOWER}}s
// @Accept json
// @Produce json
// @Success 200 {array} models.{{DOMAIN_TITLE}}Response
// @Failure 500 {object} map[string]interface{}
// @Router /{{DOMAIN_LOWER}}s [get]
func (h *{{DOMAIN_TITLE}}Handler) Get{{DOMAIN_TITLE}}s(c *gin.Context) {
	{{DOMAIN_LOWER}}s, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, {{DOMAIN_LOWER}}s)
}

// Get{{DOMAIN_TITLE}} godoc
// @Summary Get a {{DOMAIN_LOWER}} by ID
// @Description Get a single {{DOMAIN_LOWER}} by its ID
// @Tags {{DOMAIN_LOWER}}s
// @Accept json
// @Produce json
// @Param id path string true "{{DOMAIN_TITLE}} ID"
// @Success 200 {object} models.{{DOMAIN_TITLE}}Response
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /{{DOMAIN_LOWER}}s/{id} [get]
func (h *{{DOMAIN_TITLE}}Handler) Get{{DOMAIN_TITLE}}(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	{{DOMAIN_LOWER}}, err := h.service.GetByID(id)
	if err != nil {
		if err.Error() == "{{DOMAIN_LOWER}} not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, {{DOMAIN_LOWER}})
}

// Update{{DOMAIN_TITLE}} godoc
// @Summary Update a {{DOMAIN_LOWER}}
// @Description Update a {{DOMAIN_LOWER}} by its ID
// @Tags {{DOMAIN_LOWER}}s
// @Accept json
// @Produce json
// @Param id path string true "{{DOMAIN_TITLE}} ID"
// @Param {{DOMAIN_LOWER}} body models.Update{{DOMAIN_TITLE}}Request true "{{DOMAIN_TITLE}} update data"
// @Success 200 {object} models.{{DOMAIN_TITLE}}Response
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /{{DOMAIN_LOWER}}s/{id} [put]
func (h *{{DOMAIN_TITLE}}Handler) Update{{DOMAIN_TITLE}}(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req models.Update{{DOMAIN_TITLE}}Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	{{DOMAIN_LOWER}}, err := h.service.Update(id, &req)
	if err != nil {
		if err.Error() == "{{DOMAIN_LOWER}} not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, {{DOMAIN_LOWER}})
}

// Delete{{DOMAIN_TITLE}} godoc
// @Summary Delete a {{DOMAIN_LOWER}}
// @Description Delete a {{DOMAIN_LOWER}} by its ID
// @Tags {{DOMAIN_LOWER}}s
// @Accept json
// @Produce json
// @Param id path string true "{{DOMAIN_TITLE}} ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /{{DOMAIN_LOWER}}s/{id} [delete]
func (h *{{DOMAIN_TITLE}}Handler) Delete{{DOMAIN_TITLE}}(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		if err.Error() == "{{DOMAIN_LOWER}} not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}`, projectName, domainName, gcpProject)
}
