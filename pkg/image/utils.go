package image

import (
	"path/filepath"
	"strings"
)

// GenerateOutputPath generates an output filename based on input and options
func GenerateOutputPath(inputPath, suffix, extension string) string {
	dir := filepath.Dir(inputPath)
	filename := filepath.Base(inputPath)
	ext := filepath.Ext(filename)
	nameWithoutExt := strings.TrimSuffix(filename, ext)
	
	if suffix != "" {
		return filepath.Join(dir, nameWithoutExt+suffix+extension)
	}
	return filepath.Join(dir, nameWithoutExt+extension)
}

// NormalizeFormat normalizes format string to lowercase
func NormalizeFormat(format string) string {
	format = strings.ToLower(format)
	// Normalize jpeg to jpg
	if format == "jpeg" {
		return "jpg"
	}
	return format
}

// GetFileExtension returns file extension with dot
func GetFileExtension(format string) string {
	format = NormalizeFormat(format)
	return "." + format
}
