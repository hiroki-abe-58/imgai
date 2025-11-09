package image

import (
	"fmt"
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

// ConvertOptions holds options for converting an image
type ConvertOptions struct {
	Format  string
	Quality int
	Output  string
}

// ConvertImage converts an image to a different format
func ConvertImage(inputPath string, opts ConvertOptions) error {
	// Validate input file
	if err := ValidateInputFile(inputPath); err != nil {
		return err
	}

	// Normalize and validate format
	opts.Format = NormalizeFormat(opts.Format)
	if err := ValidateFormat(opts.Format); err != nil {
		return err
	}

	// Validate quality for JPEG
	if opts.Format == "jpg" {
		if err := ValidateQuality(opts.Quality); err != nil {
			return err
		}
	}

	// Open the image
	img, err := imaging.Open(inputPath)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrOpenFile, err)
	}

	// Determine output path
	outputPath := opts.Output
	if outputPath == "" {
		ext := GetFileExtension(opts.Format)
		outputPath = GenerateOutputPath(inputPath, "_converted", ext)
	}

	// Save with format-specific encoding
	if err := saveWithFormat(img, outputPath, opts.Format, opts.Quality); err != nil {
		return err
	}

	fmt.Printf("✓ Converted: %s → %s (%s)\n", inputPath, outputPath, opts.Format)
	return nil
}

// saveWithFormat saves image with specific format encoding
func saveWithFormat(img imaging.Image, outputPath, format string, quality int) error {
	switch format {
	case "jpg":
		return saveAsJPEG(img, outputPath, quality)
	case "png":
		return imaging.Save(img, outputPath)
	case "webp":
		return imaging.Save(img, outputPath)
	default:
		return fmt.Errorf("%w: %s", ErrInvalidFormat, format)
	}
}

// saveAsJPEG saves image as JPEG with specified quality
func saveAsJPEG(img imaging.Image, outputPath string, quality int) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrSaveImage, err)
	}
	defer file.Close()

	opts := &jpeg.Options{Quality: quality}
	if err := jpeg.Encode(file, img, opts); err != nil {
		return fmt.Errorf("%w: %v", ErrEncodeImage, err)
	}

	return nil
}
