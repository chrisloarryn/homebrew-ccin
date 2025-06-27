package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/chrisloarryn/chrisloarryn-cli/internal/common"
	_ "github.com/chrisloarryn/chrisloarryn-cli/internal/generators/nestjs"
	_ "github.com/chrisloarryn/chrisloarryn-cli/internal/generators/go-gin"
	_ "github.com/chrisloarryn/chrisloarryn-cli/internal/generators/go-fiber"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate CRUD applications",
	Long: `Generate complete CRUD applications with different frameworks.
Available frameworks:
- nestjs: NestJS with TypeScript
- go-gin: Go with Gin framework (REST/gRPC)
- go-fiber: Go with Fiber framework (REST/gRPC)`,
	Aliases: []string{"gen", "g"},
}

// nestjsCmd generates NestJS CRUD
var nestjsCmd = &cobra.Command{
	Use:   "nestjs [project-name]",
	Short: "Generate NestJS CRUD application",
	Long:  `Generate a complete NestJS CRUD application with TypeScript, interceptors, GCP metrics, and Docker configuration.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		domainName, _ := cmd.Flags().GetString("domain")
		gcpProject, _ := cmd.Flags().GetString("gcp-project")
		
		if domainName == "" {
			domainName = "item"
		}
		
		fmt.Printf("Generating NestJS CRUD project: %s\n", projectName)
		fmt.Printf("Domain: %s\n", domainName)
		if gcpProject != "" {
			fmt.Printf("GCP Project: %s\n", gcpProject)
		}
		
		// Get generator
		generator, err := common.Registry.Get("nestjs")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Prepare configuration
		config := &common.GeneratorConfig{
			ProjectName:  projectName,
			DomainName:   domainName,
			GCPProject:   gcpProject,
			OutputDir:    projectName,
			TemplateDir:  filepath.Join("templates", "nestjs"),
			WithGRPC:     false,
			DatabaseType: "mongodb",
			Port:         "3000",
		}

		// Generate project
		if err := generator.Generate(config); err != nil {
			fmt.Printf("Error generating project: %v\n", err)
			return
		}
		
		fmt.Printf("✅ NestJS project '%s' generated successfully!\n", projectName)
	},
}

// goGinCmd generates Go Gin CRUD
var goGinCmd = &cobra.Command{
	Use:   "go-gin [project-name]",
	Short: "Generate Go Gin CRUD application",
	Long:  `Generate a complete Go CRUD application using Gin framework with REST and gRPC support, middleware, GCP metrics, and Docker configuration.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		domainName, _ := cmd.Flags().GetString("domain")
		gcpProject, _ := cmd.Flags().GetString("gcp-project")
		grpc, _ := cmd.Flags().GetBool("grpc")
		
		if domainName == "" {
			domainName = "item"
		}
		
		fmt.Printf("Generating Go Gin CRUD project: %s\n", projectName)
		fmt.Printf("Domain: %s\n", domainName)
		if gcpProject != "" {
			fmt.Printf("GCP Project: %s\n", gcpProject)
		}
		if grpc {
			fmt.Println("Including gRPC support")
		}
		
		// Get generator
		generator, err := common.Registry.Get("go-gin")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Prepare configuration
		config := &common.GeneratorConfig{
			ProjectName:  projectName,
			DomainName:   domainName,
			GCPProject:   gcpProject,
			OutputDir:    projectName,
			TemplateDir:  filepath.Join("templates", "go-gin"),
			WithGRPC:     grpc,
			DatabaseType: "postgresql",
			Port:         "8080",
		}

		// Generate project
		if err := generator.Generate(config); err != nil {
			fmt.Printf("Error generating project: %v\n", err)
			return
		}
		
		fmt.Printf("✅ Go Gin project '%s' generated successfully!\n", projectName)
	},
}

// goFiberCmd generates Go Fiber CRUD
var goFiberCmd = &cobra.Command{
	Use:   "go-fiber [project-name]",
	Short: "Generate Go Fiber CRUD application",
	Long:  `Generate a complete Go CRUD application using Fiber framework with REST and gRPC support, middleware, GCP metrics, and Docker configuration.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		domainName, _ := cmd.Flags().GetString("domain")
		gcpProject, _ := cmd.Flags().GetString("gcp-project")
		grpc, _ := cmd.Flags().GetBool("grpc")
		
		if domainName == "" {
			domainName = "item"
		}
		
		fmt.Printf("Generating Go Fiber CRUD project: %s\n", projectName)
		fmt.Printf("Domain: %s\n", domainName)
		if gcpProject != "" {
			fmt.Printf("GCP Project: %s\n", gcpProject)
		}
		if grpc {
			fmt.Println("Including gRPC support")
		}
		
		// Get generator
		generator, err := common.Registry.Get("go-fiber")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// Prepare configuration
		config := &common.GeneratorConfig{
			ProjectName:  projectName,
			DomainName:   domainName,
			GCPProject:   gcpProject,
			OutputDir:    projectName,
			TemplateDir:  filepath.Join("templates", "go-fiber"),
			WithGRPC:     grpc,
			DatabaseType: "postgresql",
			Port:         "3000",
		}

		// Generate project
		if err := generator.Generate(config); err != nil {
			fmt.Printf("Error generating project: %v\n", err)
			return
		}
		
		fmt.Printf("✅ Go Fiber project '%s' generated successfully!\n", projectName)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(nestjsCmd)
	generateCmd.AddCommand(goGinCmd)
	generateCmd.AddCommand(goFiberCmd)

	// Add flags for all generate commands
	for _, cmd := range []*cobra.Command{nestjsCmd, goGinCmd, goFiberCmd} {
		cmd.Flags().StringP("domain", "d", "", "Domain name for the service (e.g., user, product, order)")
		cmd.Flags().StringP("gcp-project", "p", "", "GCP Project ID for metrics integration")
	}

	// Add gRPC flag for Go commands
	goGinCmd.Flags().BoolP("grpc", "g", false, "Include gRPC support")
	goFiberCmd.Flags().BoolP("grpc", "g", false, "Include gRPC support")
}
