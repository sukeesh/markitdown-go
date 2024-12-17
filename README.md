# PDF to Markdown Converter

This Go project converts PDF files to Markdown format. The program extracts text and images from the PDF, saves the images in a local `assets` folder, and generates a Markdown file referencing the images. The CLI is built using the [Cobra](https://github.com/spf13/cobra) package, providing a structured and user-friendly command-line interface.

## Features

- **Text Extraction**: Extracts text from a PDF and saves it to a Markdown file.
- **Image Extraction**: Extracts images from a PDF and stores them in an `assets` folder.
- **Image Referencing**: References the extracted images in the generated Markdown file.
- **Modular CLI**: Utilizes Cobra for a scalable and maintainable command-line interface.
- **Fully Free and Open-Source**: Built using Go libraries and CLI tools.

## Prerequisites

1. **Install Go**:
   - [Download and install Go](https://golang.org/dl/).

2. **Install the `pdfcpu` CLI tool**:
   ```bash
   go install github.com/pdfcpu/pdfcpu/cmd/pdfcpu@latest
   ```
   Ensure the `pdfcpu` executable is in your system's `PATH`.

3. **Install Project Dependencies**:
   ```bash
   go get github.com/ledongthuc/pdf
   go get github.com/spf13/cobra@latest
   ```

## Installation

1. **Clone this Repository**:
   ```bash
   git clone https://github.com/sukeesh/markitdown-go.git
   cd markitdown-go
   ```

2. **Initialize the Go Module**:
   ```bash
   go mod tidy
   ```

3. **Build the Program**:
   ```bash
   go build -o pdf2md main.go
   ```

## Usage

The CLI tool `pdf2md` provides a `convert` subcommand to handle PDF to Markdown conversion. Below are detailed instructions on how to use the tool.

### Basic Conversion

Convert a PDF file to Markdown with default settings:

```bash
./pdf2md convert <path/to/input.pdf>
```

**Example**:
```bash
./pdf2md convert sample.pdf
```

This command will:
- Generate an `output.md` file containing the extracted text and image references.
- Save extracted images in the default `assets/` directory.

### Advanced Usage

Specify custom output Markdown file and assets directory using flags:

```bash
./pdf2md convert <path/to/input.pdf> -o <path/to/output.md> -a <path/to/assets/>
```

**Example**:
```bash
./pdf2md convert sample.pdf -o output/sample.md -a output/assets/
```

### Available Flags

- `-o, --output`: Specifies the output Markdown file.
   - **Default**: `output.md`
   - **Usage**:
     ```bash
     -o path/to/output.md
     ```

- `-a, --assets`: Specifies the directory to save extracted images.
   - **Default**: `assets`
   - **Usage**:
     ```bash
     -a path/to/assets/
     ```

### Help and Documentation

Get help on the CLI usage:

```bash
./pdf2md --help
```

Get help on the `convert` subcommand:

```bash
./pdf2md convert --help
```

## Example Output

### Markdown File (`output.md`):
```markdown
# Sample PDF Title

This is some sample text extracted from the PDF.

![Image 1](assets/image_1.png)

More text on the next line.

![Image 2](assets/image_2.jpg)
```

### Assets Folder (`assets/`):
- `image_1.png`
- `image_2.jpg`

## Project Structure

```
markitdown-go/
├── cmd/
│   ├── convert.go       # Defines the 'convert' subcommand
│   └── root.go          # Defines the root command
├── pkg/
│   └── pdfconverter/
│       └── pdfconverter.go  # Core logic for PDF to Markdown conversion
├── go.mod               # Go module file
├── go.sum               # Go checksum file
├── main.go              # Entry point of the application
├── README.md            # Project documentation
└── LICENSE              # License information
```

- **`cmd/`**: Contains Cobra command definitions.
   - **`root.go`**: Defines the base command and global flags.
   - **`convert.go`**: Implements the `convert` subcommand.

- **`pkg/pdfconverter/`**: Encapsulates the core functionality for converting PDFs to Markdown.
   - **`pdfconverter.go`**: Handles text and image extraction, and Markdown generation.

- **`main.go`**: Executes the root command.

## Code Structure

- **`extractText`**: Extracts text from the PDF using the [`github.com/ledongthuc/pdf`](https://github.com/ledongthuc/pdf) library.
- **`extractImages`**: Uses the `pdfcpu` CLI tool to extract images and saves them in the specified `assets` directory.
- **`generateMarkdown`**: Combines the extracted text and image references into a Markdown file.
- **CLI Commands**: Utilizes Cobra to handle command-line arguments and flags, providing a structured interface for users.

## Dependencies

- [ledongthuc/pdf](https://github.com/ledongthuc/pdf): For text extraction.
- [pdfcpu](https://github.com/pdfcpu/pdfcpu): For image extraction.
- [Cobra](https://github.com/spf13/cobra): For building the CLI interface.

## Contributing

Contributions are welcome! Feel free to fork this repository, make changes, and submit a pull request. Whether it's improving documentation, adding new features, or fixing bugs, your contributions help make this project better.

### How to Contribute

1. **Fork the Repository**:
   Click the "Fork" button at the top-right of this repository's page.

2. **Clone Your Fork**:
   ```bash
   git clone https://github.com/yourusername/markitdown-go.git
   cd markitdown-go
   ```

3. **Create a New Branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```

4. **Make Your Changes**:
   Implement your feature or fix.

5. **Commit Your Changes**:
   ```bash
   git commit -m "Add feature: your feature description"
   ```

6. **Push to Your Fork**:
   ```bash
   git push origin feature/your-feature-name
   ```

7. **Submit a Pull Request**:
   Go to the original repository and submit a pull request from your fork.

## License

This project is licensed under the [MIT License](LICENSE). See the [LICENSE](LICENSE) file for details.
