package batch

import (
	"fmt"
	"path/filepath"
	"sync"
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
	workers int
}

// NewProcessor creates a new batch processor
func NewProcessor(workers int) *Processor {
	if workers <= 0 {
		workers = 4 // Default to 4 workers
	}
	return &Processor{workers: workers}
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

	// Create channels for work distribution
	jobs := make(chan string, len(files))
	results := make(chan Result, len(files))

	// Start worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < p.workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range jobs {
				err := processFunc(path)
				results <- Result{
					Path:    path,
					Success: err == nil,
					Error:   err,
				}
			}
		}()
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
			// Check if it's a file and hasn't been added yet
			if !seen[match] {
				files = append(files, match)
				seen[match] = true
			}
		}
	}

	return files, nil
}
