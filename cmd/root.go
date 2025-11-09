package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use:   "imgai",
	Short: "AI-powered image processing CLI tool",
	Long: `🎨 imgai - AI-powered image processing CLI tool

imgai provides modern image processing capabilities including:
  • Image resizing and optimization
  • Format conversion (PNG/JPEG/WebP)
  • Metadata management
  • AI-powered image analysis

More features coming soon!`,
	Version: version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can be added here
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.imgai.yaml)")
}
