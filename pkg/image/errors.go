package image

import "errors"

// Error definitions for image processing
var (
	// ErrInvalidFormat is returned when an unsupported format is specified
	ErrInvalidFormat = errors.New("invalid or unsupported image format")
	
	// ErrInvalidQuality is returned when quality is out of range
	ErrInvalidQuality = errors.New("quality must be between 1 and 100")
	
	// ErrInvalidDimensions is returned when dimensions are invalid
	ErrInvalidDimensions = errors.New("at least one dimension (width or height) must be specified")
	
	// ErrFileNotFound is returned when input file does not exist
	ErrFileNotFound = errors.New("input file not found")
	
	// ErrOpenFile is returned when file cannot be opened
	ErrOpenFile = errors.New("failed to open image file")
	
	// ErrDecodeImage is returned when image cannot be decoded
	ErrDecodeImage = errors.New("failed to decode image")
	
	// ErrEncodeImage is returned when image cannot be encoded
	ErrEncodeImage = errors.New("failed to encode image")
	
	// ErrSaveImage is returned when image cannot be saved
	ErrSaveImage = errors.New("failed to save image")
)
