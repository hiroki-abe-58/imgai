package cmd

import (
	"fmt"

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

Examples:
  imgai convert photo.jpg --format png
  imgai convert *.jpg --format png --dry-run
  imgai convert *.jpg --format png --workers 8`,
	Args: cobra.MinimumNArgs(1),
	RunE: runConvert,
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringVarP(&convertFormat, "format", "f", "", "Target format (jpg, png, webp) [required]")
	convertCmd.Flags().IntVarP(&convertQuality, "quality", "q", 90, "JPEG quality (1-100)")
	convertCmd.Flags().StringVarP(&convertOutput, "output", "o", "", "Output file path (single file only)")
	convertCmd.Flags().IntVar(&convertWorkers, "workers", 4, "Number of parallel workers")
	convertCmd.Flags().BoolVar(&convertDryRun, "dry-run", false, "Preview operations without executing")
	
	convertCmd.MarkFlagRequired("format")
}

func runConvert(cmd *cobra.Command, args []string) error {
	// Validate format
	convertFormat = image.NormalizeFormat(convertFormat)
	if err := image.ValidateFormat(convertFormat); err != nil {
		return err
	}

	// Validate quality for JPEG
	if convertFormat == "jpg" {
		if err := image.ValidateQuality(convertQuality); err != nil {
			return err
		}
	}

	// Dry-run mode
	if convertDryRun {
		return runConvertDryRun(args)
	}

	// Single file mode with output path
	if len(args) == 1 && convertOutput != "" {
		return runConvertSingle(args[0])
	}

	// Batch processing mode
	return runConvertBatch(args)
}

func runConvertDryRun(args []string) error {
	printDryRunHeader()
	
	processor := batch.NewProcessor(convertWorkers)
	processor.SetProgressBar(false)
	
	previewFunc := func(path string) error {
		outputPath := convertOutput
		if outputPath == "" {
			ext := image.GetFileExtension(convertFormat)
			outputPath = fmt.Sprintf("%s (auto-generated %s)", path, ext)
		}
		qualityInfo := ""
		if convertFormat == "jpg" {
			qualityInfo = fmt.Sprintf(", quality=%d", convertQuality)
		}
		fmt.Printf("  Would convert: %s → %s (%s%s)\n", path, outputPath, convertFormat, qualityInfo)
		return nil
	}
	
	results := processor.Process(args, previewFunc)
	printDryRunFooter(len(results))
	return nil
}

func runConvertSingle(inputPath string) error {
	if err := image.ValidateInputFile(inputPath); err != nil {
		return err
	}

	opts := image.ConvertOptions{
		Format:  convertFormat,
		Quality: convertQuality,
		Output:  convertOutput,
	}
	return image.ConvertImage(inputPath, opts)
}

func runConvertBatch(args []string) error {
	processor := batch.NewProcessor(convertWorkers)
	
	processFunc := func(path string) error {
		opts := image.ConvertOptions{
			Format:  convertFormat,
			Quality: convertQuality,
			Output:  "",
		}
		return image.ConvertImage(path, opts)
	}

	results := processor.Process(args, processFunc)
	return printResults(results)
}
