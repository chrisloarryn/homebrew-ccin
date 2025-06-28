package gogin

import (
	"fmt"

	"github.com/chrisloarryn/ccin/internal/common"
)

// Generator implements the Go Gin generator
type Generator struct {
	*common.BaseGenerator
}

// NewGenerator creates a new Go Gin generator
func NewGenerator() *Generator {
	return &Generator{
		BaseGenerator: common.NewBaseGenerator(
			"go-gin",
			"Generate Go CRUD application with Gin framework, REST/gRPC support, and GCP integration",
		),
	}
}

// Generate generates a Go Gin project
func (g *Generator) Generate(config *common.GeneratorConfig) error {
	// Set defaults for Go Gin
	if config.Port == "" {
		config.Port = "8080"
	}
	if config.DatabaseType == "" {
		config.DatabaseType = "postgresql"
	}

	// Prepare template data
	data := common.PrepareTemplateData(config)

	// Create template processor
	processor := common.NewTemplateProcessor(config.TemplateDir, config.OutputDir)

	// Process templates
	if err := processor.ProcessDirectory(data); err != nil {
		return fmt.Errorf("failed to process Go Gin templates: %w", err)
	}

	return nil
}

// init registers the generator
func init() {
	common.Registry.Register(NewGenerator())
}
