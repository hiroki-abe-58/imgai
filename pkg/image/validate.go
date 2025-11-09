package image

import (
	"fmt"
	"os"
	"strings"
)

// ValidateQuality checks if quality is within valid range
func ValidateQuality(quality int) error {
	if quality < MinQuality || quality > MaxQuality {
		return fmt.Errorf("%w: got %d", ErrInvalidQuality, quality)
	}
	return nil
}

// ValidateFormat checks if format is supported
func ValidateFormat(format string) error {
	format = strings.ToLower(format)
	for _, supported := range SupportedFormats {
		if format == supported {
			return nil
		}
	}
	return fmt.Errorf("%w: %s (supported: %v)", ErrInvalidFormat, format, SupportedFormats)
}

// ValidateDimensions checks if at least one dimension is specified
func ValidateDimensions(width, height int) error {
	if width <= 0 && height <= 0 {
		return ErrInvalidDimensions
	}
	return nil
}

// ValidateInputFile checks if input file exists and is readable
func ValidateInputFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrFileNotFound, path)
	}
	return nil
}
