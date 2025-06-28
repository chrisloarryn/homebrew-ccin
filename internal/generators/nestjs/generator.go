package nestjs

import (
	"fmt"

	"github.com/chrisloarryn/ccin/internal/common"
)

// Generator implements the NestJS generator
type Generator struct {
	*common.BaseGenerator
}

// NewGenerator creates a new NestJS generator
func NewGenerator() *Generator {
	return &Generator{
		BaseGenerator: common.NewBaseGenerator(
			"nestjs",
			"Generate NestJS CRUD application with TypeScript, MongoDB, and GCP integration",
		),
	}
}

// Generate generates a NestJS project
func (g *Generator) Generate(config *common.GeneratorConfig) error {
	// Set defaults for NestJS
	if config.Port == "" {
		config.Port = "3000"
	}
	if config.DatabaseType == "" {
		config.DatabaseType = "mongodb"
	}

	// Prepare template data
	data := common.PrepareTemplateData(config)

	// Create template processor
	processor := common.NewTemplateProcessor(config.TemplateDir, config.OutputDir)

	// Process templates
	if err := processor.ProcessDirectory(data); err != nil {
		return fmt.Errorf("failed to process NestJS templates: %w", err)
	}

	return nil
}

// init registers the generator
func init() {
	common.Registry.Register(NewGenerator())
}
