package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/chrisloarryn/ccin/internal/common"
	_ "github.com/chrisloarryn/ccin/internal/generators/nestjs"
	_ "github.com/chrisloarryn/ccin/internal/generators/go-gin"
	_ "github.com/chrisloarryn/ccin/internal/generators/go-fiber"
	"github.com/spf13/cobra"
)

// Constants for repeated strings to reduce duplication
const (
	defaultDomain = "item"
	
	// Flag names
	flagDomain = "domain"
	flagGCPProject = "gcp-project"
	flagGRPC = "grpc"
	
	// Generator names
	generatorNestJS = "nestjs"
	generatorGoGin = "go-gin"
	generatorGoFiber = "go-fiber"
	
	// Error messages
	errorGeneratorNotFound = "❌ Generator Error: %v\n"
	errorGeneration = "❌ Generation Error: %v\n"
	errorInvalidProjectName = "❌ Invalid project name: %v\n"
	
	// Help messages
	helpAvailableGenerators = "💡 Available generators: nestjs, go-gin, go-fiber"
	helpCheckTemplates = "💡 Check that all template files exist and are accessible"
	
	// Info messages
	msgProcessingTemplates = "📝 Processing templates..."
	nextStepsHeader = "\n🎯 Next steps:"
	readmeNote = "\n📚 Check the README.md for complete documentation"
	whatYouGetHeader = "🎯 What you'll get:\n"
	exampleHeader = "📋 Example: "
	domainLabel = "📊 Domain: "
	gcpProjectLabel = "☁️  GCP Project: "
	grpcEnabledMsg = "🔗 gRPC support enabled"
	cdCommand = "   cd %s\n"
	
	// Visual elements
	separatorLine = "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "🎯 Generate production-ready CRUD applications",
	Long: color.New(color.FgCyan, color.Bold).Sprint("🚀 GENERATE COMMAND") + color.New(color.FgWhite).Sprint(" - Create complete CRUD applications\n\n") +
		color.New(color.FgGreen).Sprint("🎯 Available Frameworks:\n") +
		color.New(color.FgYellow).Sprint("   📦 nestjs") + color.New(color.FgHiBlack).Sprint("   - NestJS with TypeScript, MongoDB, Swagger, Jest\n") +
		color.New(color.FgYellow).Sprint("   🟢 go-gin") + color.New(color.FgHiBlack).Sprint("  - Go with Gin framework, PostgreSQL, REST/gRPC\n") +
		color.New(color.FgYellow).Sprint("   ⚡ go-fiber") + color.New(color.FgHiBlack).Sprint(" - Go with Fiber framework (ultra-fast), PostgreSQL, REST/gRPC\n\n") +
		color.New(color.FgMagenta).Sprint("💡 Examples:\n") +
		color.New(color.FgHiBlack).Sprint("   ccin generate nestjs my-api --domain user --gcp-project my-project\n") +
		color.New(color.FgHiBlack).Sprint("   ccin generate go-gin orders-api --domain order --grpc\n") +
		color.New(color.FgHiBlack).Sprint("   ccin generate go-fiber products-api --domain product --gcp-project prod\n\n") +
		color.New(color.FgCyan).Sprint("🔧 Use: ") + color.New(color.FgWhite, color.Bold).Sprint("ccin generate <framework> <project-name> [flags]"),
	Aliases: []string{"gen", "g"},
}

// nestjsCmd generates NestJS CRUD
var nestjsCmd = &cobra.Command{
	Use:   "nestjs [project-name]",
	Short: "📦 Generate NestJS CRUD application",
	Long: color.New(color.FgMagenta, color.Bold).Sprint("📦 NESTJS GENERATOR\n\n") +
		color.New(color.FgGreen).Sprint(whatYouGetHeader) +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("NestJS") + color.New(color.FgHiBlack).Sprint(" framework with TypeScript\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("MongoDB") + color.New(color.FgHiBlack).Sprint(" with Mongoose ODM\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("Swagger/OpenAPI") + color.New(color.FgHiBlack).Sprint(" automatic documentation\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("GCP Metrics") + color.New(color.FgHiBlack).Sprint(" interceptors (optional)\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("Docker") + color.New(color.FgHiBlack).Sprint(" multi-stage production build\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("Jest") + color.New(color.FgHiBlack).Sprint(" testing configuration\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("ESLint + Prettier") + color.New(color.FgHiBlack).Sprint(" code quality\n\n") +
		color.New(color.FgCyan).Sprint(exampleHeader) + color.New(color.FgWhite, color.Bold).Sprint("ccin generate nestjs my-api --domain user --gcp-project my-project"),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		
		// Validate project name
		if err := validateProjectName(projectName); err != nil {
			color.New(color.FgRed, color.Bold).Printf(errorInvalidProjectName, err)
			color.New(color.FgYellow).Println("💡 Use a descriptive name like 'my-api', 'user-service', etc.")
			return
		}
		
		domainName, _ := cmd.Flags().GetString(flagDomain)
		gcpProject, _ := cmd.Flags().GetString(flagGCPProject)
		
		if domainName == "" {
			domainName = defaultDomain
		}
		
		// Print header
		printProjectHeader("NestJS", projectName, domainName, gcpProject, false)
		
		// Get generator
		generator, err := common.Registry.Get(generatorNestJS)
		if err != nil {
			handleGeneratorError(generatorNestJS, err)
			return
		}

		// Prepare configuration
		config := &common.GeneratorConfig{
			ProjectName:  projectName,
			DomainName:   domainName,
			GCPProject:   gcpProject,
			OutputDir:    projectName,
			TemplateDir:  filepath.Join("templates", generatorNestJS),
			WithGRPC:     false,
			DatabaseType: "mongodb",
			Port:         "3000",
		}

		// Generate project
		color.New(color.FgBlue).Println(msgProcessingTemplates)
		if err := generator.Generate(config); err != nil {
			handleGenerationError(err)
			return
		}
		
		// Success message
		printSuccessMessage("NestJS", projectName, []string{"npm install", "npm run start:dev"})
	},
}

// goGinCmd generates Go Gin CRUD
var goGinCmd = &cobra.Command{
	Use:   "go-gin [project-name]",
	Short: "🟢 Generate Go Gin CRUD application",
	Long: color.New(color.FgGreen, color.Bold).Sprint("🟢 GO GIN GENERATOR\n\n") +
		color.New(color.FgGreen).Sprint(whatYouGetHeader) +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("Go 1.24.4") + color.New(color.FgHiBlack).Sprint(" with Gin framework\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("PostgreSQL") + color.New(color.FgHiBlack).Sprint(" with GORM\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("REST API") + color.New(color.FgHiBlack).Sprint(" with JSON responses\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("gRPC") + color.New(color.FgHiBlack).Sprint(" support (optional with --grpc)\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("Clean Architecture") + color.New(color.FgHiBlack).Sprint(" layers\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("GCP Metrics") + color.New(color.FgHiBlack).Sprint(" middleware (optional)\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("Docker") + color.New(color.FgHiBlack).Sprint(" multi-stage production build\n\n") +
		color.New(color.FgCyan).Sprint(exampleHeader) + color.New(color.FgWhite, color.Bold).Sprint("ccin generate go-gin orders-api --domain order --grpc"),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		
		// Validate project name
		if err := validateProjectName(projectName); err != nil {
			color.New(color.FgRed, color.Bold).Printf(errorInvalidProjectName, err)
			color.New(color.FgYellow).Println("💡 Use a descriptive name like 'orders-api', 'inventory-service', etc.")
			return
		}
		
		domainName, _ := cmd.Flags().GetString(flagDomain)
		gcpProject, _ := cmd.Flags().GetString(flagGCPProject)
		grpc, _ := cmd.Flags().GetBool(flagGRPC)
		
		if domainName == "" {
			domainName = defaultDomain
		}
		
		// Print header
		printProjectHeader("Go Gin", projectName, domainName, gcpProject, grpc)
		
		// Get generator
		generator, err := common.Registry.Get(generatorGoGin)
		if err != nil {
			handleGeneratorError(generatorGoGin, err)
			return
		}

		// Prepare configuration
		config := &common.GeneratorConfig{
			ProjectName:  projectName,
			DomainName:   domainName,
			GCPProject:   gcpProject,
			OutputDir:    projectName,
			TemplateDir:  filepath.Join("templates", generatorGoGin),
			WithGRPC:     grpc,
			DatabaseType: "postgresql",
			Port:         "8080",
		}

		// Generate project
		color.New(color.FgBlue).Println(msgProcessingTemplates)
		if err := generator.Generate(config); err != nil {
			handleGenerationError(err)
			return
		}
		
		// Success message
		printSuccessMessage("Go Gin", projectName, []string{"go mod tidy", "make dev"})
	},
}

// goFiberCmd generates Go Fiber CRUD
var goFiberCmd = &cobra.Command{
	Use:   "go-fiber [project-name]",
	Short: "⚡ Generate Go Fiber CRUD application",
	Long: color.New(color.FgYellow, color.Bold).Sprint("⚡ GO FIBER GENERATOR\n\n") +
		color.New(color.FgGreen).Sprint(whatYouGetHeader) +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("Go 1.24.4") + color.New(color.FgHiBlack).Sprint(" with Fiber framework (ultra-fast!)\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("PostgreSQL") + color.New(color.FgHiBlack).Sprint(" with GORM\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("REST API") + color.New(color.FgHiBlack).Sprint(" with lightning-fast JSON responses\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("gRPC") + color.New(color.FgHiBlack).Sprint(" support (optional with --grpc)\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("Clean Architecture") + color.New(color.FgHiBlack).Sprint(" layers\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("CORS") + color.New(color.FgHiBlack).Sprint(" middleware included\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("Docker") + color.New(color.FgHiBlack).Sprint(" multi-stage production build\n\n") +
		color.New(color.FgCyan).Sprint(exampleHeader) + color.New(color.FgWhite, color.Bold).Sprint("ccin generate go-fiber products-api --domain product --gcp-project prod"),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		
		// Validate project name
		if err := validateProjectName(projectName); err != nil {
			color.New(color.FgRed, color.Bold).Printf(errorInvalidProjectName, err)
			color.New(color.FgYellow).Println("💡 Use a descriptive name like 'products-api', 'notification-service', etc.")
			return
		}
		
		domainName, _ := cmd.Flags().GetString(flagDomain)
		gcpProject, _ := cmd.Flags().GetString(flagGCPProject)
		grpc, _ := cmd.Flags().GetBool(flagGRPC)
		
		if domainName == "" {
			domainName = defaultDomain
		}
		
		// Print header
		printProjectHeader("Go Fiber", projectName, domainName, gcpProject, grpc)
		
		// Get generator
		generator, err := common.Registry.Get(generatorGoFiber)
		if err != nil {
			handleGeneratorError(generatorGoFiber, err)
			return
		}

		// Prepare configuration
		config := &common.GeneratorConfig{
			ProjectName:  projectName,
			DomainName:   domainName,
			GCPProject:   gcpProject,
			OutputDir:    projectName,
			TemplateDir:  filepath.Join("templates", generatorGoFiber),
			WithGRPC:     grpc,
			DatabaseType: "postgresql",
			Port:         "3000",
		}

		// Generate project
		color.New(color.FgBlue).Println(msgProcessingTemplates)
		if err := generator.Generate(config); err != nil {
			handleGenerationError(err)
			return
		}
		
		// Success message
		printSuccessMessage("Go Fiber", projectName, []string{"go mod tidy", "make dev"})
	},
}

// Helper functions for validation and common operations
func validateProjectName(name string) error {
	if name == "" {
		return fmt.Errorf("project name cannot be empty")
	}
	if len(name) < 2 {
		return fmt.Errorf("project name must be at least 2 characters long")
	}
	// Add more validation rules as needed
	return nil
}

func printProjectHeader(framework, projectName, domain, gcpProject string, grpc bool) {
	color.New(color.FgCyan, color.Bold).Printf("\n🚀 Generating %s CRUD project: ", framework)
	color.New(color.FgWhite, color.Bold).Printf("%s\n", projectName)
	color.New(color.FgYellow).Printf(domainLabel)
	color.New(color.FgWhite).Printf("%s\n", domain)
	if gcpProject != "" {
		color.New(color.FgMagenta).Printf(gcpProjectLabel)
		color.New(color.FgWhite).Printf("%s\n", gcpProject)
	}
	if grpc {
		color.New(color.FgBlue).Println(grpcEnabledMsg)
	}
	color.New(color.FgHiBlack).Println(separatorLine)
}

func handleGeneratorError(generatorName string, err error) {
	color.New(color.FgRed, color.Bold).Printf(errorGeneratorNotFound, err)
	color.New(color.FgYellow).Println(helpAvailableGenerators)
	color.New(color.FgHiBlack).Printf("🔍 Looking for generator: %s\n", generatorName)
}

func handleGenerationError(err error) {
	color.New(color.FgRed, color.Bold).Printf(errorGeneration, err)
	color.New(color.FgYellow).Println(helpCheckTemplates)
	color.New(color.FgHiBlack).Println("🔧 Make sure you're running from the correct directory")
}

func printSuccessMessage(framework, projectName string, commands []string) {
	color.New(color.FgGreen, color.Bold).Printf("\n✅ %s project '%s' generated successfully!\n", framework, projectName)
	color.New(color.FgCyan).Println(nextStepsHeader)
	color.New(color.FgWhite).Printf(cdCommand, projectName)
	for _, cmd := range commands {
		color.New(color.FgWhite).Printf("   %s\n", cmd)
	}
	color.New(color.FgHiBlack).Println(readmeNote)
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.AddCommand(nestjsCmd)
	generateCmd.AddCommand(goGinCmd)
	generateCmd.AddCommand(goFiberCmd)

	// Add flags for all generate commands
	for _, cmd := range []*cobra.Command{nestjsCmd, goGinCmd, goFiberCmd} {
		cmd.Flags().StringP(flagDomain, "d", "", "Domain name for the service (e.g., user, product, order)")
		cmd.Flags().StringP(flagGCPProject, "p", "", "GCP Project ID for metrics integration")
	}

	// Add gRPC flag for Go commands
	goGinCmd.Flags().BoolP(flagGRPC, "g", false, "Include gRPC support")
	goFiberCmd.Flags().BoolP(flagGRPC, "g", false, "Include gRPC support")
}
