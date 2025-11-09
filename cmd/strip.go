package cmd

import (
	"fmt"

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

Warning: By default, this command overwrites the original file.
Use --output to save to a different location.

Examples:
  imgai strip photo.jpg
  imgai strip *.jpg --dry-run
  imgai strip *.jpg --workers 8`,
	Args: cobra.MinimumNArgs(1),
	RunE: runStrip,
}

func init() {
	rootCmd.AddCommand(stripCmd)

	stripCmd.Flags().StringVarP(&stripOutput, "output", "o", "", "Output file path (single file only, default: overwrite)")
	stripCmd.Flags().IntVar(&stripWorkers, "workers", 4, "Number of parallel workers")
	stripCmd.Flags().BoolVar(&stripDryRun, "dry-run", false, "Preview operations without executing")
}

func runStrip(cmd *cobra.Command, args []string) error {
	// Dry-run mode
	if stripDryRun {
		return runStripDryRun(args)
	}

	// Single file mode with output path
	if len(args) == 1 && stripOutput != "" {
		return runStripSingle(args[0])
	}

	// Batch processing mode
	return runStripBatch(args)
}

func runStripDryRun(args []string) error {
	printDryRunHeader()
	
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
	printDryRunFooter(len(results))
	return nil
}

func runStripSingle(inputPath string) error {
	opts := metadata.StripOptions{
		Output: stripOutput,
	}
	return metadata.StripExif(inputPath, opts)
}

func runStripBatch(args []string) error {
	processor := batch.NewProcessor(stripWorkers)
	
	processFunc := func(path string) error {
		opts := metadata.StripOptions{
			Output: "",
		}
		return metadata.StripExif(path, opts)
	}

	results := processor.Process(args, processFunc)
	return printResults(results)
}
