package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ledongthuc/pdf"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <input.pdf>\n", os.Args[0])
	}

	inputPath := os.Args[1]
	outputMarkdown := "output.md"
	assetsDir := "assets"

	// Create assets directory if it doesn't exist
	if _, err := os.Stat(assetsDir); os.IsNotExist(err) {
		err := os.Mkdir(assetsDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create assets directory: %s\n", err)
		}
	}

	// Extract Text
	text, err := extractText(inputPath)
	if err != nil {
		log.Fatalf("Failed to extract text: %s\n", err)
	}

	// Extract Images
	images, err := extractImages(inputPath, assetsDir)
	if err != nil {
		log.Fatalf("Failed to extract images: %s\n", err)
	}

	// Generate Markdown
	markdown, err := generateMarkdown(text, images, assetsDir)
	if err != nil {
		log.Fatalf("Failed to generate markdown: %s\n", err)
	}

	// Write Markdown to file
	err = ioutil.WriteFile(outputMarkdown, []byte(markdown), 0644)
	if err != nil {
		log.Fatalf("Failed to write markdown file: %s\n", err)
	}

	fmt.Printf("Conversion completed. Markdown saved to %s with images in %s/\n", outputMarkdown, assetsDir)
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
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to extract images using pdfcpu: %v", err)
	}

	// Collect paths of extracted images
	files, err := ioutil.ReadDir(assetsDir)
	if err != nil {
		return nil, err
	}

	var imagePaths []string
	for _, file := range files {
		if !file.IsDir() {
			imagePaths = append(imagePaths, filepath.Join(assetsDir, file.Name()))
		}
	}

	return imagePaths, nil
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
