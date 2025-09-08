package rustaxum

import (
	"fmt"

	"github.com/chrisloarryn/ccin/internal/common"
)

// Generator implements the Rust Axum generator
type Generator struct {
	*common.BaseGenerator
}

// NewGenerator creates a new Rust Axum generator
func NewGenerator() *Generator {
	return &Generator{
		BaseGenerator: common.NewBaseGenerator(
			"rust-axum",
			"Generate Rust backend with Axum (REST) + optional Tonic (gRPC), Clean Architecture",
		),
	}
}

// Generate generates a Rust Axum project
func (g *Generator) Generate(config *common.GeneratorConfig) error {
	// Set defaults for Rust Axum
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
		return fmt.Errorf("failed to process Rust Axum templates: %w", err)
	}

	return nil
}

// init registers the generator
func init() {
	common.Registry.Register(NewGenerator())
}
