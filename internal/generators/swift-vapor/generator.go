package swiftvapor

import (
	"fmt"

	"github.com/chrisloarryn/ccin/internal/common"
)

// Generator implements the Swift Vapor generator
type Generator struct {
	*common.BaseGenerator
}

// NewGenerator creates a new Swift Vapor generator
func NewGenerator() *Generator {
	return &Generator{
		BaseGenerator: common.NewBaseGenerator(
			"swift-vapor",
			"Generate Swift Vapor backend (REST + optional gRPC) with Clean Architecture",
		),
	}
}

// Generate generates a Swift Vapor project
func (g *Generator) Generate(config *common.GeneratorConfig) error {
	// Set defaults for Swift Vapor
	if config.Port == "" {
		config.Port = "8080"
	}
	if config.DatabaseType == "" {
		config.DatabaseType = "none"
	}

	// Prepare template data
	data := common.PrepareTemplateData(config)

	// Create template processor
	processor := common.NewTemplateProcessor(config.TemplateDir, config.OutputDir)

	// Process templates
	if err := processor.ProcessDirectory(data); err != nil {
		return fmt.Errorf("failed to process Swift Vapor templates: %w", err)
	}

	return nil
}

// init registers the generator
func init() {
	common.Registry.Register(NewGenerator())
}
