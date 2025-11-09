package batch

// Config holds configuration for batch processor
type Config struct {
	Workers      int
	ShowProgress bool
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Workers:      4,
		ShowProgress: true,
	}
}

// NewConfig creates a new configuration with custom worker count
func NewConfig(workers int) *Config {
	cfg := DefaultConfig()
	if workers > 0 {
		cfg.Workers = workers
	}
	return cfg
}
