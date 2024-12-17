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

// ConvertPDFToMarkdown orchestrates the conversion from PDF to Markdown
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

// extractText extracts text from the PDF using ledongthuc/pdf
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

// extractImages extracts images from the PDF using pdfcpu CLI
func extractImages(pdfPath, assetsDir string) ([]string, error) {
	// Use pdfcpu CLI to extract images
	cmd := exec.Command("pdfcpu", "extract", "-mode", "image", pdfPath, assetsDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to extract images using pdfcpu: %v\nOutput: %s", err, string(output))
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

// generateMarkdown creates a Markdown string with text and image references
func generateMarkdown(text string, images []string, assetsDir string) (string, error) {
	var md strings.Builder

	// Write the extracted text
	md.WriteString(text)
	md.WriteString("\n\n")

	// Insert images at the end
	for i, imgPath := range images {
		relativePath := filepath.Join(assetsDir, filepath.Base(imgPath))
		md.WriteString(fmt.Sprintf("![Image %d](%s)\n\n", i+1, relativePath))
	}

	return md.String(), nil
}
