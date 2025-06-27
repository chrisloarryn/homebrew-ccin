package config

import (
	"os"
)

// Config holds application configuration
type Config struct {
	Port        string
	GRPCPort    string
	DatabaseURL string
	GinMode     string
	{{- if .GCPProject}}
	GCPProject  string
	{{- end}}
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "{{.Port}}"),
		GRPCPort:    getEnv("GRPC_PORT", "50051"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://localhost/{{.ProjectName}}_dev?sslmode=disable"),
		GinMode:     getEnv("GIN_MODE", "debug"),
		{{- if .GCPProject}}
		GCPProject:  getEnv("GCP_PROJECT", "{{.GCPProject}}"),
		{{- end}}
	}
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
