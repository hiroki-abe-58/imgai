package image

import (
	"fmt"

	"github.com/disintegration/imaging"
)

// ResizeOptions holds options for resizing an image
type ResizeOptions struct {
	Width  int
	Height int
	Output string
}

// ResizeImage resizes an image based on the provided options
func ResizeImage(inputPath string, opts ResizeOptions) error {
	// Validate input file
	if err := ValidateInputFile(inputPath); err != nil {
		return err
	}

	// Validate dimensions
	if err := ValidateDimensions(opts.Width, opts.Height); err != nil {
		return err
	}

	// Open the image
	img, err := imaging.Open(inputPath)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrOpenFile, err)
	}

	// Get original dimensions
	bounds := img.Bounds()
	origWidth := bounds.Dx()
	origHeight := bounds.Dy()

	// Calculate target dimensions
	targetWidth, targetHeight := calculateDimensions(origWidth, origHeight, opts.Width, opts.Height)

	// Resize the image
	resized := imaging.Resize(img, targetWidth, targetHeight, imaging.Lanczos)

	// Determine output path
	outputPath := opts.Output
	if outputPath == "" {
		suffix := fmt.Sprintf("_resized_%dx%d", targetWidth, targetHeight)
		outputPath = GenerateOutputPath(inputPath, suffix, ".jpg")
	}

	// Save the resized image
	if err := imaging.Save(resized, outputPath); err != nil {
		return fmt.Errorf("%w: %v", ErrSaveImage, err)
	}

	fmt.Printf("✓ Resized: %s → %s (%dx%d)\n", inputPath, outputPath, targetWidth, targetHeight)
	return nil
}

// calculateDimensions calculates target dimensions while maintaining aspect ratio
func calculateDimensions(origWidth, origHeight, targetWidth, targetHeight int) (int, int) {
	if targetWidth > 0 && targetHeight > 0 {
		// Both dimensions specified, use as-is
		return targetWidth, targetHeight
	}

	aspectRatio := float64(origWidth) / float64(origHeight)

	if targetWidth > 0 {
		// Only width specified, calculate height
		return targetWidth, int(float64(targetWidth) / aspectRatio)
	}

	// Only height specified, calculate width
	return int(float64(targetHeight) * aspectRatio), targetHeight
}
