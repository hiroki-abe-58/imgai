# 🎨 imgai

AI-powered image processing CLI tool for modern workflows.

## Features

### ✅ Implemented

#### Core Image Processing
- 🖼️ **Image resizing** - Maintain aspect ratio or exact dimensions
- 🔄 **Format conversion** - Convert between JPEG, PNG, and WebP
- 📊 **Batch processing** - Process multiple images in parallel with goroutines
- 🎯 **Quality optimization** - Adjust JPEG quality for optimal file size
- 📈 **Progress bar** - Visual feedback for batch operations

#### Metadata Management
- 📝 **EXIF reading** - View camera settings, GPS, and image metadata

### 🔜 Planned

#### Metadata Management (Phase 2)
- 🔒 Privacy-focused metadata removal
- 📋 Bulk metadata operations

#### AI-Powered Features (Phase 3)
- 🤖 Image description generation
- 📖 OCR (Optical Character Recognition)
- 🎯 Object detection
- 🔍 Smart image analysis

## Installation

### Prerequisites
- Go 1.21 or higher
- macOS (Apple Silicon optimized) / Linux / Windows

### From Source
```bash
git clone https://github.com/hiroki-abe-58/imgai.git
cd imgai
go build -o imgai
```

## Usage

### Resize Images
```bash
# Resize single image to 800px width (maintain aspect ratio)
imgai resize input.jpg --width 800

# Resize to 600px height
imgai resize input.jpg --height 600

# Resize to exact dimensions
imgai resize input.jpg --width 800 --height 600

# Batch resize with progress bar
imgai resize *.jpg --width 800
```

### Convert Formats
```bash
# Convert to PNG
imgai convert input.jpg --format png

# Convert to JPEG with custom quality
imgai convert input.png --format jpg --quality 85

# Batch convert all PNGs to WebP
imgai convert *.png --format webp
```

### View EXIF Data
```bash
# Display EXIF metadata
imgai exif photo.jpg
```

### Advanced Options
```bash
# Use 8 parallel workers for faster batch processing
imgai resize *.jpg --width 800 --workers 8

# Specify output file (single file only)
imgai resize input.jpg --width 800 --output result.jpg
```

## Development

### Building
```bash
go build -o imgai
```

### Testing
```bash
go test ./...
```

## Roadmap

- [x] Project initialization
- [x] CLI framework setup (cobra)
- [x] Basic image operations (resize, convert)
- [x] Batch processing with goroutines
- [x] EXIF metadata reading
- [x] Progress bar for batch operations
- [ ] EXIF metadata removal
- [ ] Dry-run mode
- [ ] AI integration
- [ ] Cross-platform packaging

## Architecture
```
imgai/
├── cmd/           # CLI commands
│   ├── root.go    # Root command
│   ├── resize.go  # Resize command
│   ├── convert.go # Convert command
│   └── exif.go    # EXIF command
├── pkg/           # Core packages
│   ├── image/     # Image processing logic
│   ├── batch/     # Batch processing with progress
│   └── metadata/  # EXIF metadata handling
└── main.go        # Entry point
```

## License

MIT License - see [LICENSE](LICENSE) for details

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Author

Developed by [hiroki-abe-58](https://github.com/hiroki-abe-58)
