package i18n

import (
	"os"
)

var currentLang = "en"
var messages map[string]map[string]string

// Init initializes the i18n system
func Init(lang string) error {
	if lang == "" {
		// Try to get from environment
		lang = os.Getenv("IMGAI_LANG")
		if lang == "" {
			lang = "en" // Default to English
		}
	}
	
	currentLang = lang
	return loadMessages()
}

// loadMessages loads messages from embedded data
func loadMessages() error {
	messages = make(map[string]map[string]string)
	
	// English messages
	messages["en"] = map[string]string{
		// Common
		"app_description": "AI-powered image processing CLI tool",
		"dry_run_mode": "🔍 DRY RUN MODE - No files will be modified",
		"would_process": "✓ Would process %d images",
		"run_without_dry_run": "💡 Run without --dry-run to execute",
		"successfully_processed": "✓ Successfully processed %d/%d images",
		"failed": "✗ Failed: %s - %v",
		"file_not_found": "file not found: %s",
		
		// Resize
		"resize_short": "Resize one or multiple images",
		"would_resize": "Would resize: %s → %s (%dx%d)",
		"resized": "✓ Resized: %s → %s (%dx%d)",
		"dimension_required": "at least one dimension (width or height) must be specified",
		
		// Convert
		"convert_short": "Convert one or multiple images to a different format",
		"would_convert": "Would convert: %s → %s (%s%s)",
		"converted": "✓ Converted: %s → %s (%s)",
		"quality_range": "quality must be between 1 and 100",
		
		// EXIF
		"exif_short": "Display EXIF metadata from an image",
		"exif_data_for": "EXIF Data for: %s",
		"no_exif": "No EXIF data found in this image.",
		
		// Strip
		"strip_short": "Remove EXIF metadata from images",
		"would_strip": "Would strip metadata: %s → %s",
		"stripped": "✓ Stripped metadata: %s",
	}
	
	// Japanese messages
	messages["ja"] = map[string]string{
		// Common
		"app_description": "AI搭載の画像処理CLIツール",
		"dry_run_mode": "🔍 ドライランモード - ファイルは変更されません",
		"would_process": "✓ %d個の画像を処理する予定です",
		"run_without_dry_run": "💡 --dry-run なしで実行すると処理が実行されます",
		"successfully_processed": "✓ %d/%d個の画像を正常に処理しました",
		"failed": "✗ 失敗: %s - %v",
		"file_not_found": "ファイルが見つかりません: %s",
		
		// Resize
		"resize_short": "1つまたは複数の画像をリサイズ",
		"would_resize": "リサイズ予定: %s → %s (%dx%d)",
		"resized": "✓ リサイズ完了: %s → %s (%dx%d)",
		"dimension_required": "幅または高さのいずれかを指定する必要があります",
		
		// Convert
		"convert_short": "1つまたは複数の画像を別の形式に変換",
		"would_convert": "変換予定: %s → %s (%s%s)",
		"converted": "✓ 変換完了: %s → %s (%s)",
		"quality_range": "品質は1から100の間で指定してください",
		
		// EXIF
		"exif_short": "画像のEXIFメタデータを表示",
		"exif_data_for": "EXIFデータ: %s",
		"no_exif": "この画像にはEXIFデータが見つかりませんでした。",
		
		// Strip
		"strip_short": "画像からEXIFメタデータを削除",
		"would_strip": "メタデータ削除予定: %s → %s",
		"stripped": "✓ メタデータ削除完了: %s",
	}
	
	return nil
}

// T translates a message key
func T(key string) string {
	if messages == nil {
		Init("")
	}
	
	if msg, ok := messages[currentLang][key]; ok {
		return msg
	}
	
	// Fallback to English
	if msg, ok := messages["en"][key]; ok {
		return msg
	}
	
	return key
}

// GetLang returns the current language
func GetLang() string {
	return currentLang
}
