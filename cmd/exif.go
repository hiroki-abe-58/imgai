package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/hiroki-abe-58/imgai/pkg/metadata"
	"github.com/spf13/cobra"
)

var exifCmd = &cobra.Command{
	Use:   "exif [image]",
	Short: "Display EXIF metadata from an image",
	Long: `Display EXIF metadata information from an image file.

EXIF (Exchangeable Image File Format) contains metadata such as:
  • Camera make and model
  • Date and time taken
  • Camera settings (ISO, aperture, shutter speed, focal length)
  • Image dimensions
  • GPS coordinates (if available)
  • Orientation

Examples:
  # Display EXIF data
  imgai exif photo.jpg

  # View GPS coordinates
  imgai exif IMG_1234.jpg`,
	Args: cobra.ExactArgs(1),
	RunE: runExif,
}

func init() {
	rootCmd.AddCommand(exifCmd)
}

func runExif(cmd *cobra.Command, args []string) error {
	inputPath := args[0]

	// Validate input file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", inputPath)
	}

	// Read EXIF data
	data, err := metadata.ReadExif(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read EXIF data: %w", err)
	}

	// Format and display
	fmt.Printf("EXIF Data for: %s\n", inputPath)
	fmt.Println(strings.Repeat("-", 50))
	output := metadata.FormatExif(data)
	if output == "" {
		fmt.Println("No EXIF data found in this image.")
	} else {
		fmt.Print(output)
	}

	return nil
}
