package gofiber

import (
	"fmt"

	"github.com/chrisloarryn/chrisloarryn-cli/internal/common"
)

// Generator implements the Go Fiber generator
type Generator struct {
	*common.BaseGenerator
}

// NewGenerator creates a new Go Fiber generator
func NewGenerator() *Generator {
	return &Generator{
		BaseGenerator: common.NewBaseGenerator(
			"go-fiber",
			"Generate Go CRUD application with Fiber framework, REST/gRPC support, and GCP integration",
		),
	}
}

// Generate generates a Go Fiber project
func (g *Generator) Generate(config *common.GeneratorConfig) error {
	// Set defaults for Go Fiber
	if config.Port == "" {
		config.Port = "3000"
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
		return fmt.Errorf("failed to process Go Fiber templates: %w", err)
	}

	return nil
}

// init registers the generator
func init() {
	common.Registry.Register(NewGenerator())
}
