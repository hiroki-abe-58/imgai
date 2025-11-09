package cmd

import (
	"fmt"
	"os"

	"github.com/hiroki-abe-58/imgai/pkg/image"
	"github.com/spf13/cobra"
)

var (
	convertFormat  string
	convertQuality int
	convertOutput  string
)

var convertCmd = &cobra.Command{
	Use:   "convert [image]",
	Short: "Convert an image to a different format",
	Long: `Convert an image to a different format (JPEG, PNG, WebP).

Supported formats:
  • JPEG/JPG - Lossy compression, best for photos
  • PNG - Lossless compression, best for graphics
  • WebP - Modern format with superior compression

Examples:
  # Convert to PNG
  imgai convert input.jpg --format png

  # Convert to JPEG with custom quality
  imgai convert input.png --format jpg --quality 85

  # Convert to WebP
  imgai convert input.jpg --format webp

  # Specify output file
  imgai convert input.jpg --format png --output result.png`,
	Args: cobra.ExactArgs(1),
	RunE: runConvert,
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringVarP(&convertFormat, "format", "f", "", "Target format (jpg, png, webp) [required]")
	convertCmd.Flags().IntVarP(&convertQuality, "quality", "q", 90, "JPEG quality (1-100, default: 90)")
	convertCmd.Flags().StringVarP(&convertOutput, "output", "o", "", "Output file path (default: input.format)")
	
	convertCmd.MarkFlagRequired("format")
}

func runConvert(cmd *cobra.Command, args []string) error {
	inputPath := args[0]

	// Validate input file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", inputPath)
	}

	// Validate quality range
	if convertQuality < 1 || convertQuality > 100 {
		return fmt.Errorf("quality must be between 1 and 100")
	}

	// Prepare convert options
	opts := image.ConvertOptions{
		Format:  convertFormat,
		Quality: convertQuality,
		Output:  convertOutput,
	}

	// Convert the image
	return image.ConvertImage(inputPath, opts)
}
