package main

import (
	"fmt"
	"log"
	"os"

	"{{.ProjectName}}/internal/api"
	"{{.ProjectName}}/internal/config"
	"{{.ProjectName}}/internal/database"
	{{- if .WithGRPC}}
	"{{.ProjectName}}/internal/grpc"
	{{- end}}
	{{- if .GCPProject}}
	"{{.ProjectName}}/internal/metrics"
	{{- end}}
	
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg := config.Load()

	{{- if .GCPProject}}
	// Initialize metrics
	if err := metrics.Initialize(cfg.GCPProject); err != nil {
		log.Fatalf("Failed to initialize metrics: %v", err)
	}
	{{- end}}

	// Initialize database
	db, err := database.Initialize(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize Gin router
	gin.SetMode(cfg.GinMode)
	router := gin.Default()

	// Initialize API routes
	api.SetupRoutes(router, db)

	{{- if .WithGRPC}}
	// Start gRPC server in a goroutine
	go func() {
		if err := grpc.StartServer(cfg.GRPCPort, db); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()
	{{- end}}

	// Start HTTP server
	port := cfg.Port
	if port == "" {
		port = "{{.Port}}"
	}

	fmt.Printf("🚀 Server starting on port %s\n", port)
	{{- if .WithGRPC}}
	fmt.Printf("🚀 gRPC server starting on port %s\n", cfg.GRPCPort)
	{{- end}}
	
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
