package cmd

import (
	"fmt"

	"github.com/hiroki-abe-58/imgai/pkg/batch"
	"github.com/hiroki-abe-58/imgai/pkg/image"
	"github.com/spf13/cobra"
)

var (
	resizeWidth   int
	resizeHeight  int
	resizeOutput  string
	resizeWorkers int
	resizeDryRun  bool
)

var resizeCmd = &cobra.Command{
	Use:   "resize [image(s)]",
	Short: "Resize one or multiple images",
	Long: `Resize one or multiple images to specified dimensions.

If only width or height is specified, the aspect ratio will be maintained.
If both are specified, images will be resized to exact dimensions.

Examples:
  imgai resize photo.jpg --width 800
  imgai resize *.jpg --width 800 --dry-run
  imgai resize *.jpg --width 800 --workers 8`,
	Args: cobra.MinimumNArgs(1),
	RunE: runResize,
}

func init() {
	rootCmd.AddCommand(resizeCmd)

	resizeCmd.Flags().IntVarP(&resizeWidth, "width", "w", 0, "Target width in pixels")
	resizeCmd.Flags().IntVar(&resizeHeight, "height", 0, "Target height in pixels")
	resizeCmd.Flags().StringVarP(&resizeOutput, "output", "o", "", "Output file path (single file only)")
	resizeCmd.Flags().IntVar(&resizeWorkers, "workers", 4, "Number of parallel workers")
	resizeCmd.Flags().BoolVar(&resizeDryRun, "dry-run", false, "Preview operations without executing")
}

func runResize(cmd *cobra.Command, args []string) error {
	// Validate dimensions
	if err := image.ValidateDimensions(resizeWidth, resizeHeight); err != nil {
		return err
	}

	// Dry-run mode
	if resizeDryRun {
		return runResizeDryRun(args)
	}

	// Single file mode with output path
	if len(args) == 1 && resizeOutput != "" {
		return runResizeSingle(args[0])
	}

	// Batch processing mode
	return runResizeBatch(args)
}

func runResizeDryRun(args []string) error {
	printDryRunHeader()
	
	processor := batch.NewProcessor(resizeWorkers)
	processor.SetProgressBar(false)
	
	previewFunc := func(path string) error {
		outputPath := resizeOutput
		if outputPath == "" {
			outputPath = fmt.Sprintf("%s (auto-generated)", path)
		}
		fmt.Printf("  Would resize: %s → %s (%dx%d)\n", path, outputPath, resizeWidth, resizeHeight)
		return nil
	}
	
	results := processor.Process(args, previewFunc)
	printDryRunFooter(len(results))
	return nil
}

func runResizeSingle(inputPath string) error {
	if err := image.ValidateInputFile(inputPath); err != nil {
		return err
	}

	opts := image.ResizeOptions{
		Width:  resizeWidth,
		Height: resizeHeight,
		Output: resizeOutput,
	}
	return image.ResizeImage(inputPath, opts)
}

func runResizeBatch(args []string) error {
	processor := batch.NewProcessor(resizeWorkers)
	
	processFunc := func(path string) error {
		opts := image.ResizeOptions{
			Width:  resizeWidth,
			Height: resizeHeight,
			Output: "",
		}
		return image.ResizeImage(path, opts)
	}

	results := processor.Process(args, processFunc)
	return printResults(results)
}
