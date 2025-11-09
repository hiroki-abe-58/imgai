package metadata

import (
	"fmt"

	"github.com/disintegration/imaging"
)

// StripExif removes all EXIF metadata from an image
func StripExif(inputPath string, opts StripOptions) error {
	// Open the image
	img, err := imaging.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}

	// Determine output path
	outputPath := getOutputPath(inputPath, opts.Output)

	// Save the image without metadata
	if err := imaging.Save(img, outputPath); err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	fmt.Printf("✓ Stripped metadata: %s\n", outputPath)
	return nil
}

// getOutputPath returns the appropriate output path
func getOutputPath(inputPath, customOutput string) string {
	if customOutput != "" {
		return customOutput
	}
	return inputPath // Overwrite original
}
