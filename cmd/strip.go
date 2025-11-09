package cmd

import (
	"fmt"
	"os"

	"github.com/hiroki-abe-58/imgai/pkg/batch"
	"github.com/hiroki-abe-58/imgai/pkg/metadata"
	"github.com/spf13/cobra"
)

var (
	stripOutput  string
	stripWorkers int
	stripDryRun  bool
)

var stripCmd = &cobra.Command{
	Use:   "strip [image(s)]",
	Short: "Remove EXIF metadata from images",
	Long: `Remove all EXIF metadata from one or multiple images for privacy protection.

This command removes:
  • Camera make and model
  • Date and time information
  • GPS coordinates
  • Camera settings
  • All other metadata

Warning: By default, this command overwrites the original file.
Use --output to save to a different location.

Examples:
  # Strip metadata from single image (overwrites original)
  imgai strip photo.jpg

  # Strip metadata and save to new file
  imgai strip photo.jpg --output clean.jpg

  # Strip metadata from multiple images
  imgai strip img1.jpg img2.jpg img3.jpg

  # Preview without actually stripping (dry-run)
  imgai strip *.jpg --dry-run

  # Batch strip with 8 parallel workers
  imgai strip *.jpg --workers 8`,
	Args: cobra.MinimumNArgs(1),
	RunE: runStrip,
}

func init() {
	rootCmd.AddCommand(stripCmd)

	stripCmd.Flags().StringVarP(&stripOutput, "output", "o", "", "Output file path (single file only, default: overwrite)")
	stripCmd.Flags().IntVar(&stripWorkers, "workers", 4, "Number of parallel workers for batch processing")
	stripCmd.Flags().BoolVar(&stripDryRun, "dry-run", false, "Preview operations without executing them")
}

func runStrip(cmd *cobra.Command, args []string) error {
	// Dry-run mode
	if stripDryRun {
		fmt.Println("🔍 DRY RUN MODE - No files will be modified")
		fmt.Println()
		
		processor := batch.NewProcessor(stripWorkers)
		processor.SetProgressBar(false)
		
		previewFunc := func(path string) error {
			outputPath := stripOutput
			if outputPath == "" {
				outputPath = path + " (overwrite)"
			}
			fmt.Printf("  Would strip metadata: %s → %s\n", path, outputPath)
			return nil
		}
		
		results := processor.Process(args, previewFunc)
		fmt.Printf("\n✓ Would process %d images\n", len(results))
		fmt.Println("💡 Run without --dry-run to execute")
		return nil
	}

	// Single file mode with output path
	if len(args) == 1 && stripOutput != "" {
		inputPath := args[0]
		if _, err := os.Stat(inputPath); os.IsNotExist(err) {
			return fmt.Errorf("file not found: %s", inputPath)
		}

		opts := metadata.StripOptions{
			Output: stripOutput,
		}
		return metadata.StripExif(inputPath, opts)
	}

	// Batch processing mode
	processor := batch.NewProcessor(stripWorkers)
	
	processFunc := func(path string) error {
		opts := metadata.StripOptions{
			Output: "", // Overwrite original
		}
		return metadata.StripExif(path, opts)
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
