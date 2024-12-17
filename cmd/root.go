package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pdf2md",
	Short: "Convert PDF files to Markdown format",
	Long: `pdf2md is a CLI tool that extracts text and images from PDF files
and generates corresponding Markdown (.md) files with embedded images.`,
	// Uncomment the following line if your bare application has an action associated with it
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you can define persistent flags and configuration settings
	rootCmd.PersistentFlags().StringP("output", "o", "output.md", "Output Markdown file")
	rootCmd.PersistentFlags().StringP("assets", "a", "assets", "Directory to save extracted images")

	// Add subcommands
	rootCmd.AddCommand(convertCmd)
}
