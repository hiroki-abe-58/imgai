package cmd

import (
	"fmt"
	"os"

	"github.com/hiroki-abe-58/imgai/pkg/batch"
)

// printDryRunHeader prints the dry-run mode header
func printDryRunHeader() {
	fmt.Println("🔍 DRY RUN MODE - No files will be modified")
	fmt.Println()
}

// printDryRunFooter prints the dry-run mode footer
func printDryRunFooter(count int) {
	fmt.Printf("\n✓ Would process %d images\n", count)
	fmt.Println("💡 Run without --dry-run to execute")
}

// printResults prints processing results summary
func printResults(results []batch.Result) error {
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
