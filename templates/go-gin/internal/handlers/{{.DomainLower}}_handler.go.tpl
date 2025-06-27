package handlers

import (
	"net/http"
	"strconv"

	"{{.ProjectName}}/internal/models"
	"{{.ProjectName}}/internal/services"

	"github.com/gin-gonic/gin"
)

// {{.DomainTitle}}Handler handles HTTP requests for {{.DomainLower}}s
type {{.DomainTitle}}Handler struct {
	service *services.{{.DomainTitle}}Service
}

// New{{.DomainTitle}}Handler creates a new {{.DomainLower}} handler
func New{{.DomainTitle}}Handler(service *services.{{.DomainTitle}}Service) *{{.DomainTitle}}Handler {
	return &{{.DomainTitle}}Handler{service: service}
}

// GetAll handles GET /{{.DomainLower}}s
func (h *{{.DomainTitle}}Handler) GetAll(c *gin.Context) {
	{{.DomainLower}}s, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": {{.DomainLower}}s})
}

// GetByID handles GET /{{.DomainLower}}s/:id
func (h *{{.DomainTitle}}Handler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	{{.DomainLower}}, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": {{.DomainLower}}})
}

// Create handles POST /{{.DomainLower}}s
func (h *{{.DomainTitle}}Handler) Create(c *gin.Context) {
	var req models.Create{{.DomainTitle}}Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	{{.DomainLower}}, err := h.service.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": {{.DomainLower}}})
}

// Update handles PUT /{{.DomainLower}}s/:id
func (h *{{.DomainTitle}}Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req models.Update{{.DomainTitle}}Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	{{.DomainLower}}, err := h.service.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": {{.DomainLower}}})
}

// Delete handles DELETE /{{.DomainLower}}s/:id
func (h *{{.DomainTitle}}Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "{{.DomainTitle}} deleted successfully"})
}
