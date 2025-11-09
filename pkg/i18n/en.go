package i18n

// getEnglishMessages returns English language messages
func getEnglishMessages() map[string]string {
	return map[string]string{
		// Common
		"app_description":        "AI-powered image processing CLI tool",
		"dry_run_mode":           "🔍 DRY RUN MODE - No files will be modified",
		"would_process":          "✓ Would process %d images",
		"run_without_dry_run":    "💡 Run without --dry-run to execute",
		"successfully_processed": "✓ Successfully processed %d/%d images",
		"failed":                 "✗ Failed: %s - %v",
		"file_not_found":         "file not found: %s",

		// Resize
		"resize_short":        "Resize one or multiple images",
		"would_resize":        "Would resize: %s → %s (%dx%d)",
		"resized":             "✓ Resized: %s → %s (%dx%d)",
		"dimension_required":  "at least one dimension (width or height) must be specified",

		// Convert
		"convert_short":   "Convert one or multiple images to a different format",
		"would_convert":   "Would convert: %s → %s (%s%s)",
		"converted":       "✓ Converted: %s → %s (%s)",
		"quality_range":   "quality must be between 1 and 100",

		// EXIF
		"exif_short":    "Display EXIF metadata from an image",
		"exif_data_for": "EXIF Data for: %s",
		"no_exif":       "No EXIF data found in this image.",

		// Strip
		"strip_short":  "Remove EXIF metadata from images",
		"would_strip":  "Would strip metadata: %s → %s",
		"stripped":     "✓ Stripped metadata: %s",
	}
}
