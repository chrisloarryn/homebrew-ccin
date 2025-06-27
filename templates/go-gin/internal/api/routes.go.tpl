package api

import (
	"database/sql"

	"{{.ProjectName}}/internal/handlers"
	"{{.ProjectName}}/internal/middleware"
	"{{.ProjectName}}/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	{{- if .GCPProject}}
	router.Use(middleware.MetricsMiddleware())
	{{- end}}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// {{.DomainTitle}} routes
		{{.DomainLower}}Service := services.New{{.DomainTitle}}Service(db)
		{{.DomainLower}}Handler := handlers.New{{.DomainTitle}}Handler({{.DomainLower}}Service)

		{{.DomainLower}}Routes := v1.Group("/{{.DomainLower}}s")
		{
			{{.DomainLower}}Routes.GET("", {{.DomainLower}}Handler.GetAll)
			{{.DomainLower}}Routes.GET("/:id", {{.DomainLower}}Handler.GetByID)
			{{.DomainLower}}Routes.POST("", {{.DomainLower}}Handler.Create)
			{{.DomainLower}}Routes.PUT("/:id", {{.DomainLower}}Handler.Update)
			{{.DomainLower}}Routes.DELETE("/:id", {{.DomainLower}}Handler.Delete)
		}
	}
}
