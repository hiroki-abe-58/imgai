package cmd

import (
	"fmt"
	"os"

	"github.com/hiroki-abe-58/imgai/pkg/image"
	"github.com/spf13/cobra"
)

var (
	resizeWidth  int
	resizeHeight int
	resizeOutput string
)

var resizeCmd = &cobra.Command{
	Use:   "resize [image]",
	Short: "Resize an image",
	Long: `Resize an image to specified dimensions.

If only width or height is specified, the aspect ratio will be maintained.
If both are specified, the image will be resized to exact dimensions.

Examples:
  # Resize to 800px width (maintain aspect ratio)
  imgai resize input.jpg --width 800

  # Resize to 600px height (maintain aspect ratio)
  imgai resize input.jpg --height 600

  # Resize to exact dimensions
  imgai resize input.jpg --width 800 --height 600

  # Specify output file
  imgai resize input.jpg --width 800 --output output.jpg`,
	Args: cobra.ExactArgs(1),
	RunE: runResize,
}

func init() {
	rootCmd.AddCommand(resizeCmd)

	resizeCmd.Flags().IntVarP(&resizeWidth, "width", "w", 0, "Target width in pixels")
	resizeCmd.Flags().IntVar(&resizeHeight, "height", 0, "Target height in pixels")
	resizeCmd.Flags().StringVarP(&resizeOutput, "output", "o", "", "Output file path (default: input_resized_WxH.ext)")
}

func runResize(cmd *cobra.Command, args []string) error {
	inputPath := args[0]

	// Validate input file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", inputPath)
	}

	// Validate at least one dimension is specified
	if resizeWidth == 0 && resizeHeight == 0 {
		return fmt.Errorf("at least one dimension (width or height) must be specified")
	}

	// Prepare resize options
	opts := image.ResizeOptions{
		Width:  resizeWidth,
		Height: resizeHeight,
		Output: resizeOutput,
	}

	// Resize the image
	return image.ResizeImage(inputPath, opts)
}
