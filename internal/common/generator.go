package common

import (
	"strings"
)

// Generator represents a code generator
type Generator interface {
	Generate(config *GeneratorConfig) error
	GetName() string
	GetDescription() string
}

// GeneratorConfig holds configuration for generators
type GeneratorConfig struct {
	ProjectName  string
	DomainName   string
	GCPProject   string
	OutputDir    string
	TemplateDir  string
	WithGRPC     bool
	DatabaseType string
	Port         string
}

// PrepareTemplateData prepares data for template processing
func PrepareTemplateData(config *GeneratorConfig) *TemplateData {
	return &TemplateData{
		ProjectName:  config.ProjectName,
		DomainName:   config.DomainName,
		DomainTitle:  strings.Title(config.DomainName),
		DomainUpper:  strings.ToUpper(config.DomainName),
		DomainLower:  strings.ToLower(config.DomainName),
		GCPProject:   config.GCPProject,
		WithGRPC:     config.WithGRPC,
		Port:         config.Port,
		DatabaseType: config.DatabaseType,
	}
}

// BaseGenerator provides common functionality for all generators
type BaseGenerator struct {
	name        string
	description string
}

// NewBaseGenerator creates a new base generator
func NewBaseGenerator(name, description string) *BaseGenerator {
	return &BaseGenerator{
		name:        name,
		description: description,
	}
}

// GetName returns the generator name
func (bg *BaseGenerator) GetName() string {
	return bg.name
}

// GetDescription returns the generator description
func (bg *BaseGenerator) GetDescription() string {
	return bg.description
}
