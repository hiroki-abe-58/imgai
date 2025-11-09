package cmd

import (
	"fmt"
	"os"

	"github.com/hiroki-abe-58/imgai/pkg/i18n"
	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
)

func getLongDescription() string {
	if i18n.GetLang() == "ja" {
		return `🎨 imgai - AI搭載の画像処理CLIツール

imgaiは以下の機能を提供します：
  • 画像のリサイズと最適化
  • フォーマット変換（PNG/JPEG/WebP）
  • 並列実行によるバッチ処理
  • EXIFメタデータの読み取りと削除
  • プログレスバーとドライランモード

Goで構築され、Apple Siliconに最適化されています。

使用例:
  imgai resize photo.jpg --width 800
  imgai convert image.png --format jpg
  imgai exif photo.jpg
  imgai strip photo.jpg`
	}
	return `🎨 imgai - AI-powered image processing CLI tool

imgai provides modern image processing capabilities including:
  • Image resizing and optimization
  • Format conversion (PNG/JPEG/WebP)
  • Batch processing with parallel execution
  • EXIF metadata reading and removal
  • Progress bar and dry-run mode

Built with Go and optimized for Apple Silicon.

Examples:
  imgai resize photo.jpg --width 800
  imgai convert image.png --format jpg
  imgai exif photo.jpg
  imgai strip photo.jpg`
}

var rootCmd = &cobra.Command{
	Use:     "imgai",
	Short:   i18n.T("app_description"),
	Version: version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Update Long description based on current language
		cmd.Long = getLongDescription()
	},
}

func Execute() {
	// Set Long description before execution
	rootCmd.Long = getLongDescription()
	
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can be added here
}
