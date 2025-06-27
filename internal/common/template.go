package common

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// TemplateData represents the data passed to templates
type TemplateData struct {
	ProjectName   string
	DomainName    string
	DomainTitle   string
	DomainUpper   string
	DomainLower   string
	GCPProject    string
	WithGRPC      bool
	Port          string
	DatabaseType  string
}

// TemplateProcessor handles template processing
type TemplateProcessor struct {
	templateDir string
	outputDir   string
}

// NewTemplateProcessor creates a new template processor
func NewTemplateProcessor(templateDir, outputDir string) *TemplateProcessor {
	return &TemplateProcessor{
		templateDir: templateDir,
		outputDir:   outputDir,
	}
}

// ProcessTemplate processes a single template file
func (tp *TemplateProcessor) ProcessTemplate(templatePath, outputPath string, data *TemplateData) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		return err
	}

	// Read template file
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	// Create output file
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Execute template
	return tmpl.Execute(file, data)
}

// ProcessDirectory processes all templates in a directory recursively
func (tp *TemplateProcessor) ProcessDirectory(data *TemplateData) error {
	return filepath.Walk(tp.templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Calculate relative path
		relPath, err := filepath.Rel(tp.templateDir, path)
		if err != nil {
			return err
		}

		// Remove .tpl extension from output path
		outputPath := filepath.Join(tp.outputDir, relPath)
		if filepath.Ext(outputPath) == ".tpl" {
			outputPath = outputPath[:len(outputPath)-4]
		}

		// Replace template variables in path
		outputPath = tp.replacePlaceholders(outputPath, data)

		// Process template
		return tp.ProcessTemplate(path, outputPath, data)
	})
}

// replacePlaceholders replaces template placeholders in file paths
func (tp *TemplateProcessor) replacePlaceholders(path string, data *TemplateData) string {
	path = strings.ReplaceAll(path, "{{.DomainLower}}", data.DomainLower)
	path = strings.ReplaceAll(path, "{{.DomainTitle}}", data.DomainTitle)
	path = strings.ReplaceAll(path, "{{.DomainUpper}}", data.DomainUpper)
	path = strings.ReplaceAll(path, "{{.ProjectName}}", data.ProjectName)
	path = strings.ReplaceAll(path, "/domain/", "/"+data.DomainLower+"/")
	return path
}
