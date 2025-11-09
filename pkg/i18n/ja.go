package i18n

// getJapaneseMessages returns Japanese language messages
func getJapaneseMessages() map[string]string {
	return map[string]string{
		// Common
		"app_description":        "AI搭載の画像処理CLIツール",
		"dry_run_mode":           "🔍 ドライランモード - ファイルは変更されません",
		"would_process":          "✓ %d個の画像を処理する予定です",
		"run_without_dry_run":    "💡 --dry-run なしで実行すると処理が実行されます",
		"successfully_processed": "✓ %d/%d個の画像を正常に処理しました",
		"failed":                 "✗ 失敗: %s - %v",
		"file_not_found":         "ファイルが見つかりません: %s",

		// Resize
		"resize_short":       "1つまたは複数の画像をリサイズ",
		"would_resize":       "リサイズ予定: %s → %s (%dx%d)",
		"resized":            "✓ リサイズ完了: %s → %s (%dx%d)",
		"dimension_required": "幅または高さのいずれかを指定する必要があります",

		// Convert
		"convert_short": "1つまたは複数の画像を別の形式に変換",
		"would_convert": "変換予定: %s → %s (%s%s)",
		"converted":     "✓ 変換完了: %s → %s (%s)",
		"quality_range": "品質は1から100の間で指定してください",

		// EXIF
		"exif_short":    "画像のEXIFメタデータを表示",
		"exif_data_for": "EXIFデータ: %s",
		"no_exif":       "この画像にはEXIFデータが見つかりませんでした。",

		// Strip
		"strip_short": "画像からEXIFメタデータを削除",
		"would_strip": "メタデータ削除予定: %s → %s",
		"stripped":    "✓ メタデータ削除完了: %s",
	}
}
