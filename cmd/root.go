package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccin",
	Short: "🚀 Advanced CLI for generating modern CRUD applications",
	Long: color.New(color.FgCyan, color.Bold).Sprint("🎯 CCIN CLI") + color.New(color.FgWhite).Sprint(" - ChrisLoarryn's Comprehensive Code Integration & Initialization Tool\n\n") +
		color.New(color.FgGreen).Sprint("✨ Generate production-ready CRUD applications with multiple frameworks:\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("NestJS") + color.New(color.FgHiBlack).Sprint(" (Node.js 24.2.0 + TypeScript + MongoDB)\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("Go + Gin") + color.New(color.FgHiBlack).Sprint(" (REST/gRPC + PostgreSQL + Clean Architecture)\n") +
		color.New(color.FgYellow).Sprint("   • ") + color.New(color.FgWhite).Sprint("Go + Fiber") + color.New(color.FgHiBlack).Sprint(" (Ultra-fast REST/gRPC + PostgreSQL)\n\n") +
		color.New(color.FgMagenta).Sprint("🎁 What you get:\n") +
		color.New(color.FgHiGreen).Sprint("   ✅ Complete CRUD operations\n") +
		color.New(color.FgHiGreen).Sprint("   ✅ Production-ready Docker configuration\n") +
		color.New(color.FgHiGreen).Sprint("   ✅ GCP metrics & monitoring integration\n") +
		color.New(color.FgHiGreen).Sprint("   ✅ Automated Makefiles for all workflows\n") +
		color.New(color.FgHiGreen).Sprint("   ✅ API documentation (Swagger/OpenAPI)\n") +
		color.New(color.FgHiGreen).Sprint("   ✅ Clean Architecture patterns\n") +
		color.New(color.FgHiGreen).Sprint("   ✅ Template-based code generation\n\n") +
		color.New(color.FgCyan).Sprint("🚀 Quick start: ") + color.New(color.FgWhite, color.Bold).Sprint("ccin generate --help"),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		color.New(color.FgRed, color.Bold).Fprintf(os.Stderr, "\n❌ Command Error: %v\n", err)
		color.New(color.FgYellow).Fprint(os.Stderr, "💡 Quick help: ")
		color.New(color.FgCyan, color.Bold).Fprint(os.Stderr, "ccin --help")
		color.New(color.FgYellow).Fprint(os.Stderr, " or ")
		color.New(color.FgCyan, color.Bold).Fprint(os.Stderr, "ccin generate --help")
		color.New(color.FgHiBlack).Fprintln(os.Stderr, "\n📚 Documentation: https://github.com/chrisloarryn/homebrew-ccin")
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccin.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")
	
	// Handle version flag
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetBool("version")
		if version {
			color.New(color.FgCyan, color.Bold).Println("🎯 CCIN CLI - ChrisLoarryn's Comprehensive Code Integration & Initialization Tool")
			color.New(color.FgWhite).Println("Version: 1.0.0")
			color.New(color.FgHiBlack).Println("Author: Chris Loarryn (@chrisloarryn)")
			color.New(color.FgHiBlack).Println("Repository: https://github.com/chrisloarryn/homebrew-ccin")
			color.New(color.FgGreen).Println("\n✨ Generate production-ready CRUD applications with modern frameworks!")
			return
		}
		cmd.Help()
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ccin" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".ccin")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		color.New(color.FgHiBlack).Fprintf(os.Stderr, "📁 Using config file: %s\n", viper.ConfigFileUsed())
	}
}
