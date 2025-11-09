package cmd

import (
	"fmt"
	"os"

	"github.com/hiroki-abe-58/imgai/pkg/batch"
	"github.com/hiroki-abe-58/imgai/pkg/image"
	"github.com/spf13/cobra"
)

var (
	convertFormat  string
	convertQuality int
	convertOutput  string
	convertWorkers int
	convertDryRun  bool
)

var convertCmd = &cobra.Command{
	Use:   "convert [image(s)]",
	Short: "Convert one or multiple images to a different format",
	Long: `Convert one or multiple images to a different format (JPEG, PNG, WebP).

Supported formats:
  • JPEG/JPG - Lossy compression, best for photos
  • PNG - Lossless compression, best for graphics
  • WebP - Modern format with superior compression

Batch processing: When multiple files are provided, they will be processed
in parallel using goroutines for maximum performance.

Examples:
  # Convert single image to PNG
  imgai convert input.jpg --format png

  # Convert multiple images to JPEG
  imgai convert img1.png img2.png --format jpg --quality 85

  # Preview without actually converting (dry-run)
  imgai convert *.jpg --format png --dry-run

  # Convert with 8 parallel workers
  imgai convert *.jpg --format png --workers 8`,
	Args: cobra.MinimumNArgs(1),
	RunE: runConvert,
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringVarP(&convertFormat, "format", "f", "", "Target format (jpg, png, webp) [required]")
	convertCmd.Flags().IntVarP(&convertQuality, "quality", "q", 90, "JPEG quality (1-100, default: 90)")
	convertCmd.Flags().StringVarP(&convertOutput, "output", "o", "", "Output file path (single file only)")
	convertCmd.Flags().IntVar(&convertWorkers, "workers", 4, "Number of parallel workers for batch processing")
	convertCmd.Flags().BoolVar(&convertDryRun, "dry-run", false, "Preview operations without executing them")
	
	convertCmd.MarkFlagRequired("format")
}

func runConvert(cmd *cobra.Command, args []string) error {
	// Validate quality range
	if convertQuality < 1 || convertQuality > 100 {
		return fmt.Errorf("quality must be between 1 and 100")
	}

	// Dry-run mode
	if convertDryRun {
		fmt.Println("🔍 DRY RUN MODE - No files will be modified")
		fmt.Println()
		
		processor := batch.NewProcessor(convertWorkers)
		processor.SetProgressBar(false) // Disable progress bar in dry-run
		
		previewFunc := func(path string) error {
			opts := image.ConvertOptions{
				Format:  convertFormat,
				Quality: convertQuality,
				Output:  convertOutput,
			}
			outputPath := opts.Output
			if outputPath == "" {
				// Generate preview of output filename
				outputPath = fmt.Sprintf("%s (auto-generated .%s)", path, convertFormat)
			}
			qualityInfo := ""
			if convertFormat == "jpg" || convertFormat == "jpeg" {
				qualityInfo = fmt.Sprintf(", quality=%d", convertQuality)
			}
			fmt.Printf("  Would convert: %s → %s (%s%s)\n", path, outputPath, convertFormat, qualityInfo)
			return nil
		}
		
		results := processor.Process(args, previewFunc)
		fmt.Printf("\n✓ Would process %d images\n", len(results))
		fmt.Println("💡 Run without --dry-run to execute")
		return nil
	}

	// Single file mode with output path
	if len(args) == 1 && convertOutput != "" {
		inputPath := args[0]
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			return fmt.Errorf("file not found: %s", inputPath)
		}

		opts := image.ConvertOptions{
			Format:  convertFormat,
			Quality: convertQuality,
			Output:  convertOutput,
		}
		return image.ConvertImage(inputPath, opts)
	}

	// Batch processing mode
	processor := batch.NewProcessor(convertWorkers)
	
	processFunc := func(path string) error {
		opts := image.ConvertOptions{
			Format:  convertFormat,
			Quality: convertQuality,
			Output:  "", // Auto-generate output path
		}
		return image.ConvertImage(path, opts)
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
