package metadata

import (
	"fmt"

	"github.com/disintegration/imaging"
)

// StripOptions holds options for stripping metadata
type StripOptions struct {
	Output string
}

// StripExif removes all EXIF metadata from an image
func StripExif(inputPath string, opts StripOptions) error {
	// Open the image
	img, err := imaging.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}

	// Determine output path
	outputPath := opts.Output
	if outputPath == "" {
		outputPath = inputPath // Overwrite original
	}

	// Save the image without metadata
	// imaging.Save automatically strips EXIF data
	if err := imaging.Save(img, outputPath); err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	fmt.Printf("✓ Stripped metadata: %s\n", outputPath)
	return nil
}
