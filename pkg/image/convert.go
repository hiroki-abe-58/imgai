package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"golang.org/x/image/webp"
)

// ConvertOptions holds options for converting images
type ConvertOptions struct {
	Format  string
	Quality int
	Output  string
}

// SupportedFormats lists all supported image formats
var SupportedFormats = []string{"jpg", "jpeg", "png", "webp"}

// ConvertImage converts an image to a different format
func ConvertImage(inputPath string, opts ConvertOptions) error {
	// Validate format
	format := strings.ToLower(opts.Format)
	if !isFormatSupported(format) {
		return fmt.Errorf("unsupported format: %s (supported: %v)", format, SupportedFormats)
	}

	// Open and decode the image
	src, err := openImage(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}

	// Determine output path
	outputPath := opts.Output
	if outputPath == "" {
		outputPath = generateConvertOutputPath(inputPath, format)
	}

	// Save with the target format
	if err := saveImage(src, outputPath, format, opts.Quality); err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	fmt.Printf("✓ Converted: %s → %s (%s)\n", inputPath, outputPath, strings.ToUpper(format))
	return nil
}

// openImage opens and decodes an image file
func openImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Detect format and decode
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".webp":
		return webp.Decode(file)
	default:
		// Use imaging library's Open for other formats
		return imaging.Open(path)
	}
}

// saveImage saves an image in the specified format
func saveImage(img image.Image, path, format string, quality int) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case "jpg", "jpeg":
		if quality == 0 {
			quality = 90 // Default JPEG quality
		}
		return jpeg.Encode(file, img, &jpeg.Options{Quality: quality})
	
	case "png":
		encoder := png.Encoder{CompressionLevel: png.DefaultCompression}
		return encoder.Encode(file, img)
	
	case "webp":
		// Use imaging library for WebP encoding
		return imaging.Save(img, path)
	
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

// isFormatSupported checks if a format is supported
func isFormatSupported(format string) bool {
	format = strings.ToLower(format)
	for _, f := range SupportedFormats {
		if f == format {
			return true
		}
	}
	return false
}

// generateConvertOutputPath generates an output filename for the converted image
func generateConvertOutputPath(inputPath, targetFormat string) string {
	ext := filepath.Ext(inputPath)
	base := inputPath[:len(inputPath)-len(ext)]
	return fmt.Sprintf("%s.%s", base, targetFormat)
}
