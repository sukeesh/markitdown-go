package pdfconverter

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ledongthuc/pdf"
)

// ConvertPDFToMarkdown converts a PDF file to Markdown format, extracting both text and images.
// The images are saved to assetsDir and referenced in the resulting markdown file.
func ConvertPDFToMarkdown(pdfPath, outputMarkdown, assetsDir string) error {
	// Create assets directory if it doesn't exist
	err := createAssetsDir(assetsDir)
	if err != nil {
		return fmt.Errorf("failed to create assets directory: %v", err)
	}

	// Extract Text
	text, err := extractText(pdfPath)
	if err != nil {
		return fmt.Errorf("failed to extract text: %v", err)
	}

	// Extract Images
	images, err := extractImages(pdfPath, assetsDir)
	if err != nil {
		return fmt.Errorf("failed to extract images: %v", err)
	}

	// Generate Markdown
	markdown, err := generateMarkdown(text, images, assetsDir)
	if err != nil {
		return fmt.Errorf("failed to generate markdown: %v", err)
	}

	// Write Markdown to file
	err = os.WriteFile(outputMarkdown, []byte(markdown), 0644)
	if err != nil {
		return fmt.Errorf("failed to write markdown file: %v", err)
	}

	return nil
}

// createAssetsDir creates the assets directory if it doesn't exist
func createAssetsDir(assetsDir string) error {
	if _, err := os.Stat(assetsDir); os.IsNotExist(err) {
		err := os.Mkdir(assetsDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// extractText extracts plain text content from a PDF file using the ledongthuc/pdf library.
// It preserves line breaks and returns the text as a single string.
func extractText(pdfPath string) (string, error) {
	f, r, err := pdf.Open(pdfPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		line := scanner.Text()
		buf.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// extractImages extracts all images from the PDF using pdfcpu CLI tool and saves them
// to the specified assets directory. Returns a list of extracted image file paths.
func extractImages(pdfPath, assetsDir string) ([]string, error) {
	cmd := exec.Command("pdfcpu", "extract", "-mode", "image", pdfPath, assetsDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("pdfcpu image extraction failed: %w\nOutput: %s", err, output)
	}

	// Collect paths of extracted images
	files, err := ioutil.ReadDir(assetsDir)
	if err != nil {
		return nil, err
	}

	var imagePaths []string
	for _, file := range files {
		if !file.IsDir() && isImageFile(file.Name()) {
			imagePaths = append(imagePaths, filepath.Join(assetsDir, file.Name()))
		}
	}

	return imagePaths, nil
}

// isImageFile checks if a file has an image extension
func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif", ".bmp", ".tiff", ".svg":
		return true
	default:
		return false
	}
}

// generateMarkdown creates a Markdown document combining the extracted text and image references.
// Images are appended at the end of the document with auto-incrementing numbers.
func generateMarkdown(text string, images []string, assetsDir string) (string, error) {
	var md strings.Builder

	// Add text content
	md.WriteString(text)

	// Add a separator between text and images if there are images
	if len(images) > 0 {
		md.WriteString("\n\n---\n\n")
	}

	// Add image references
	for i, imgPath := range images {
		relativePath := filepath.Join(assetsDir, filepath.Base(imgPath))
		md.WriteString(fmt.Sprintf("![Image %d](%s)\n", i+1, relativePath))
	}

	return md.String(), nil
}
