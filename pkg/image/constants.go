package image

// Default values for image processing
const (
	// DefaultQuality is the default JPEG quality
	DefaultQuality = 90
	
	// MinQuality is the minimum allowed JPEG quality
	MinQuality = 1
	
	// MaxQuality is the maximum allowed JPEG quality
	MaxQuality = 100
	
	// DefaultFormat is the default output format
	DefaultFormat = "jpg"
)

// Supported image formats
var SupportedFormats = []string{"jpg", "jpeg", "png", "webp"}
