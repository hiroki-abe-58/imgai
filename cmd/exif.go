package cmd

import (
	"fmt"
	"strings"

	"github.com/hiroki-abe-58/imgai/pkg/image"
	"github.com/hiroki-abe-58/imgai/pkg/metadata"
	"github.com/spf13/cobra"
)

var exifCmd = &cobra.Command{
	Use:   "exif [image]",
	Short: "Display EXIF metadata from an image",
	Long: `Display EXIF metadata information from an image file.

EXIF contains metadata such as camera settings, GPS coordinates, and more.

Examples:
  imgai exif photo.jpg
  imgai exif IMG_1234.jpg`,
	Args: cobra.ExactArgs(1),
	RunE: runExif,
}

func init() {
	rootCmd.AddCommand(exifCmd)
}

func runExif(cmd *cobra.Command, args []string) error {
	inputPath := args[0]

	// Validate input file
	if err := image.ValidateInputFile(inputPath); err != nil {
		return err
	}

	// Read EXIF data
	data, err := metadata.ReadExif(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read EXIF data: %w", err)
	}

	// Display results
	displayExifData(inputPath, data)
	return nil
}

func displayExifData(path string, data *metadata.ExifData) {
	fmt.Printf("EXIF Data for: %s\n", path)
	fmt.Println(strings.Repeat("-", 50))
	
	if data.IsEmpty() {
		fmt.Println("No EXIF data found in this image.")
		return
	}
	
	output := metadata.FormatExif(data)
	fmt.Print(output)
}
