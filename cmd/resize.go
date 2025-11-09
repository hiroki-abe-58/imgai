package cmd

import (
	"fmt"
	"os"

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

Batch processing: When multiple files are provided, they will be processed
in parallel using goroutines for maximum performance.

Examples:
  # Resize single image to 800px width
  imgai resize input.jpg --width 800

  # Resize multiple images
  imgai resize img1.jpg img2.jpg img3.jpg --width 800

  # Preview without actually resizing (dry-run)
  imgai resize *.jpg --width 800 --dry-run

  # Resize with 8 parallel workers
  imgai resize *.jpg --width 800 --workers 8`,
	Args: cobra.MinimumNArgs(1),
	RunE: runResize,
}

func init() {
	rootCmd.AddCommand(resizeCmd)

	resizeCmd.Flags().IntVarP(&resizeWidth, "width", "w", 0, "Target width in pixels")
	resizeCmd.Flags().IntVar(&resizeHeight, "height", 0, "Target height in pixels")
	resizeCmd.Flags().StringVarP(&resizeOutput, "output", "o", "", "Output file path (single file only)")
	resizeCmd.Flags().IntVar(&resizeWorkers, "workers", 4, "Number of parallel workers for batch processing")
	resizeCmd.Flags().BoolVar(&resizeDryRun, "dry-run", false, "Preview operations without executing them")
}

func runResize(cmd *cobra.Command, args []string) error {
	// Validate at least one dimension is specified
	if resizeWidth == 0 && resizeHeight == 0 {
		return fmt.Errorf("at least one dimension (width or height) must be specified")
	}

	// Dry-run mode
	if resizeDryRun {
		fmt.Println("🔍 DRY RUN MODE - No files will be modified")
		fmt.Println()
		
		processor := batch.NewProcessor(resizeWorkers)
		processor.SetProgressBar(false) // Disable progress bar in dry-run
		
		previewFunc := func(path string) error {
			opts := image.ResizeOptions{
				Width:  resizeWidth,
				Height: resizeHeight,
				Output: resizeOutput,
			}
			outputPath := opts.Output
			if outputPath == "" {
				// Generate preview of output filename
				outputPath = fmt.Sprintf("%s (auto-generated)", path)
			}
			fmt.Printf("  Would resize: %s → %s (%dx%d)\n", path, outputPath, resizeWidth, resizeHeight)
			return nil
		}
		
		results := processor.Process(args, previewFunc)
		fmt.Printf("\n✓ Would process %d images\n", len(results))
		fmt.Println("💡 Run without --dry-run to execute")
		return nil
	}

	// Single file mode with output path
	if len(args) == 1 && resizeOutput != "" {
		inputPath := args[0]
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			return fmt.Errorf("file not found: %s", inputPath)
		}

		opts := image.ResizeOptions{
			Width:  resizeWidth,
			Height: resizeHeight,
			Output: resizeOutput,
		}
		return image.ResizeImage(inputPath, opts)
	}

	// Batch processing mode
	processor := batch.NewProcessor(resizeWorkers)
	
	processFunc := func(path string) error {
		opts := image.ResizeOptions{
			Width:  resizeWidth,
			Height: resizeHeight,
			Output: "", // Auto-generate output path
		}
		return image.ResizeImage(path, opts)
	}

	results := processor.Process(args, processFunc)

	// Display summary
	successCount := 0
	for _, result := range results {
		if result.Success {
			successCount++
		} else {
			fmt.Fprintf(os.Stderr, "✗ Failed: %s - %v\n", result.Path, result.Error)
		}
	}

	fmt.Printf("\n✓ Successfully processed %d/%d images\n", successCount, len(results))

	if successCount < len(results) {
		return fmt.Errorf("some images failed to process")
	}

	return nil
}
