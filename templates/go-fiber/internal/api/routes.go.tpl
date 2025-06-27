package api

import (
	"database/sql"

	"{{.ProjectName}}/internal/handlers"
	"{{.ProjectName}}/internal/middleware"
	"{{.ProjectName}}/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// SetupRoutes configures all API routes
func SetupRoutes(app *fiber.App, db *sql.DB) {
	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())
	{{- if .GCPProject}}
	app.Use(middleware.MetricsMiddleware())
	{{- end}}

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
		})
	})

	// API v1 routes
	v1 := app.Group("/api/v1")

	// {{.DomainTitle}} routes
	{{.DomainLower}}Service := services.New{{.DomainTitle}}Service(db)
	{{.DomainLower}}Handler := handlers.New{{.DomainTitle}}Handler({{.DomainLower}}Service)

	{{.DomainLower}}Routes := v1.Group("/{{.DomainLower}}s")
	{{.DomainLower}}Routes.Get("/", {{.DomainLower}}Handler.GetAll)
	{{.DomainLower}}Routes.Get("/:id", {{.DomainLower}}Handler.GetByID)
	{{.DomainLower}}Routes.Post("/", {{.DomainLower}}Handler.Create)
	{{.DomainLower}}Routes.Put("/:id", {{.DomainLower}}Handler.Update)
	{{.DomainLower}}Routes.Delete("/:id", {{.DomainLower}}Handler.Delete)
}
