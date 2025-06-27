package handlers

import (
	"strconv"

	"{{.ProjectName}}/internal/models"
	"{{.ProjectName}}/internal/services"

	"github.com/gofiber/fiber/v2"
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
func (h *{{.DomainTitle}}Handler) GetAll(c *fiber.Ctx) error {
	{{.DomainLower}}s, err := h.service.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": {{.DomainLower}}s,
	})
}

// GetByID handles GET /{{.DomainLower}}s/:id
func (h *{{.DomainTitle}}Handler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	{{.DomainLower}}, err := h.service.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": {{.DomainLower}},
	})
}

// Create handles POST /{{.DomainLower}}s
func (h *{{.DomainTitle}}Handler) Create(c *fiber.Ctx) error {
	var req models.Create{{.DomainTitle}}Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Basic validation
	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	{{.DomainLower}}, err := h.service.Create(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": {{.DomainLower}},
	})
}

// Update handles PUT /{{.DomainLower}}s/:id
func (h *{{.DomainTitle}}Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var req models.Update{{.DomainTitle}}Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	{{.DomainLower}}, err := h.service.Update(id, &req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": {{.DomainLower}},
	})
}

// Delete handles DELETE /{{.DomainLower}}s/:id
func (h *{{.DomainTitle}}Handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := h.service.Delete(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "{{.DomainTitle}} deleted successfully",
	})
}
