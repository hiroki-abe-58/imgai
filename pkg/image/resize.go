package image

import (
	"fmt"
	"image"
	"path/filepath"

	"github.com/disintegration/imaging"
)

// ResizeOptions holds options for resizing images
type ResizeOptions struct {
	Width  int
	Height int
	Output string
}

// ResizeImage resizes an image file
func ResizeImage(inputPath string, opts ResizeOptions) error {
	// Open the image
	src, err := imaging.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}

	// Calculate dimensions
	width, height := calculateDimensions(src, opts.Width, opts.Height)

	// Resize the image using Lanczos resampling filter
	resized := imaging.Resize(src, width, height, imaging.Lanczos)

	// Determine output path
	outputPath := opts.Output
	if outputPath == "" {
		outputPath = generateOutputPath(inputPath, width, height)
	}

	// Save the resized image
	if err := imaging.Save(resized, outputPath); err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	fmt.Printf("✓ Resized: %s → %s (%dx%d)\n", inputPath, outputPath, width, height)
	return nil
}

// calculateDimensions calculates the target dimensions while maintaining aspect ratio
func calculateDimensions(img image.Image, targetWidth, targetHeight int) (int, int) {
	bounds := img.Bounds()
	srcWidth := bounds.Dx()
	srcHeight := bounds.Dy()

	// If both dimensions are specified, use them
	if targetWidth > 0 && targetHeight > 0 {
		return targetWidth, targetHeight
	}

	// If only width is specified, calculate height maintaining aspect ratio
	if targetWidth > 0 && targetHeight == 0 {
		aspectRatio := float64(srcHeight) / float64(srcWidth)
		return targetWidth, int(float64(targetWidth) * aspectRatio)
	}

	// If only height is specified, calculate width maintaining aspect ratio
	if targetWidth == 0 && targetHeight > 0 {
		aspectRatio := float64(srcWidth) / float64(srcHeight)
		return int(float64(targetHeight) * aspectRatio), targetHeight
	}

	// If neither is specified, return original dimensions
	return srcWidth, srcHeight
}

// generateOutputPath generates an output filename for the resized image
func generateOutputPath(inputPath string, width, height int) string {
	ext := filepath.Ext(inputPath)
	base := inputPath[:len(inputPath)-len(ext)]
	return fmt.Sprintf("%s_resized_%dx%d%s", base, width, height, ext)
}
