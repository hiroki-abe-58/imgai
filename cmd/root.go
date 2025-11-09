package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use:   "imgai",
	Short: "AI-powered image processing CLI tool",
	Long: `🎨 imgai - AI-powered image processing CLI tool

imgai provides modern image processing capabilities including:
  • Image resizing and optimization
  • Format conversion (PNG/JPEG/WebP)
  • Batch processing with parallel execution
  • EXIF metadata reading and removal
  • Progress bar and dry-run mode

Built with Go and optimized for Apple Silicon.

Examples:
  imgai resize photo.jpg --width 800
  imgai convert image.png --format jpg
  imgai exif photo.jpg
  imgai strip photo.jpg`,
	Version: version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can be added here
}
