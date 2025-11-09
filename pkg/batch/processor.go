package batch

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/schollz/progressbar/v3"
)

// ProcessFunc is a function type for processing a single file
type ProcessFunc func(path string) error

// Result holds the result of processing a file
type Result struct {
	Path    string
	Success bool
	Error   error
}

// Processor handles batch processing of files
type Processor struct {
	config *Config
}

// NewProcessor creates a new batch processor with default config
func NewProcessor(workers int) *Processor {
	return &Processor{
		config: NewConfig(workers),
	}
}

// NewProcessorWithConfig creates a new batch processor with custom config
func NewProcessorWithConfig(config *Config) *Processor {
	if config == nil {
		config = DefaultConfig()
	}
	return &Processor{
		config: config,
	}
}

// SetProgressBar enables or disables the progress bar
func (p *Processor) SetProgressBar(show bool) {
	p.config.ShowProgress = show
}

// Process processes multiple files concurrently
func (p *Processor) Process(patterns []string, processFunc ProcessFunc) []Result {
	// Expand patterns to file paths
	files, err := expandPatterns(patterns)
	if err != nil {
		return []Result{{
			Path:    patterns[0],
			Success: false,
			Error:   fmt.Errorf("failed to expand patterns: %w", err),
		}}
	}

	if len(files) == 0 {
		return []Result{{
			Path:    patterns[0],
			Success: false,
			Error:   fmt.Errorf("no files found matching patterns"),
		}}
	}

	// Create progress bar if enabled
	var bar *progressbar.ProgressBar
	if p.config.ShowProgress && len(files) > 1 {
		bar = createProgressBar(len(files))
	}

	// Process files concurrently
	results := p.processFiles(files, processFunc, bar)

	if bar != nil {
		fmt.Println() // New line after progress bar
	}

	return results
}

// processFiles processes files using worker pool pattern
func (p *Processor) processFiles(files []string, processFunc ProcessFunc, bar *progressbar.ProgressBar) []Result {
	jobs := make(chan string, len(files))
	results := make(chan Result, len(files))

	// Start worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < p.config.Workers; i++ {
		wg.Add(1)
		go p.worker(&wg, jobs, results, processFunc, bar)
	}

	// Send jobs to workers
	for _, file := range files {
		jobs <- file
	}
	close(jobs)

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	var allResults []Result
	for result := range results {
		allResults = append(allResults, result)
	}

	return allResults
}

// worker processes jobs from the jobs channel
func (p *Processor) worker(wg *sync.WaitGroup, jobs <-chan string, results chan<- Result, processFunc ProcessFunc, bar *progressbar.ProgressBar) {
	defer wg.Done()
	for path := range jobs {
		err := processFunc(path)
		results <- Result{
			Path:    path,
			Success: err == nil,
			Error:   err,
		}
		if bar != nil {
			bar.Add(1)
		}
	}
}

// createProgressBar creates a configured progress bar
func createProgressBar(total int) *progressbar.ProgressBar {
	return progressbar.NewOptions(total,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(40),
		progressbar.OptionSetDescription("[cyan]Processing images...[reset]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
}

// expandPatterns expands glob patterns to actual file paths
func expandPatterns(patterns []string) ([]string, error) {
	var files []string
	seen := make(map[string]bool)

	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, err
		}

		for _, match := range matches {
			if !seen[match] {
				files = append(files, match)
				seen[match] = true
			}
		}
	}

	return files, nil
}
