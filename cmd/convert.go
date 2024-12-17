package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/sukeesh/markitdown-go/pkg/pdfconverter"
)

var convertCmd = &cobra.Command{
	Use:   "convert [pdf file]",
	Short: "Convert a PDF file to Markdown",
	Long: `Convert a specified PDF file to a Markdown (.md) file, extracting text and images.

Example:
  pdf2md convert sample.pdf -o result.md -a images`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]

		// Get flags
		outputMarkdown, _ := cmd.Flags().GetString("output")
		assetsDir, _ := cmd.Flags().GetString("assets")

		// Validate input file
		absInputPath, err := filepath.Abs(inputPath)
		if err != nil {
			log.Fatalf("Invalid input path: %v", err)
		}

		// Perform conversion
		err = pdfconverter.ConvertPDFToMarkdown(absInputPath, outputMarkdown, assetsDir)
		if err != nil {
			log.Fatalf("Conversion failed: %v", err)
		}

		fmt.Printf("Conversion completed. Markdown saved to %s with images in %s/\n", outputMarkdown, assetsDir)
	},
}

func init() {
	// Define flags specific to the convert command
	convertCmd.Flags().StringP("output", "o", "output.md", "Output Markdown file")
	convertCmd.Flags().StringP("assets", "a", "assets", "Directory to save extracted images")

	// Mark flags as required if necessary
	// convertCmd.MarkFlagRequired("output")
}
