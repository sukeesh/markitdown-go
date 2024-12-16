# PDF to Markdown Converter

This Go project converts PDF files to Markdown format. The program extracts text and images from the PDF, saves the images in a local `assets` folder, and generates a Markdown file referencing the images.

## Features

- Extracts text from a PDF and saves it to a Markdown file.
- Extracts images from a PDF and stores them in an `assets` folder.
- References the extracted images in the generated Markdown file.
- Fully free and open-source, using Go libraries and CLI tools.

## Prerequisites

1. Install Go:
    - [Download and install Go](https://golang.org/dl/).

2. Install the `pdfcpu` CLI tool:
   ```bash
   go install github.com/pdfcpu/pdfcpu/cmd/pdfcpu@latest
   ```
   Ensure the `pdfcpu` executable is in your system's `PATH`.

3. Install project dependencies:
   ```bash
   go get github.com/ledongthuc/pdf
   ```

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/sukeesh/markitdown-go.git
   cd pdf-to-markdown
   ```

2. Initialize the Go module:
   ```bash
   go mod init pdf-to-markdown
   ```

3. Build the program:
   ```bash
   go build -o pdf2md main.go
   ```

## Usage

Run the program with the path to a PDF file as an argument:

```bash
./pdf2md <path/to/input.pdf>
```

Example:
```bash
./pdf2md sample.pdf
```

### Output

1. `output.md`: The generated Markdown file containing the text and references to images.
2. `assets/`: A folder containing the extracted images (e.g., `image_1.png`, `image_2.jpg`).

### Example Output

#### Markdown File (`output.md`):
```markdown
# Sample PDF Title

This is some sample text extracted from the PDF.

![Image 1](assets/image_1.png)

More text on the next line.

![Image 2](assets/image_2.jpg)
```

#### Assets Folder (`assets/`):
- `image_1.png`
- `image_2.jpg`

## Code Structure

- **`extractText`**: Extracts text from the PDF using the `github.com/ledongthuc/pdf` library.
- **`extractImages`**: Uses the `pdfcpu` CLI tool to extract images and saves them in the `assets` directory.
- **`generateMarkdown`**: Combines the extracted text and image references into a Markdown file.

## Dependencies

- [ledongthuc/pdf](https://github.com/ledongthuc/pdf): For text extraction.
- [pdfcpu](https://github.com/pdfcpu/pdfcpu): For image extraction.

## Contributing

Contributions are welcome! Feel free to fork this repository, make changes, and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

