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

// loadMessages loads messages from language-specific files
func loadMessages() error {
	messages = make(map[string]map[string]string)
	
	// Load English messages
	messages["en"] = getEnglishMessages()
	
	// Load Japanese messages
	messages["ja"] = getJapaneseMessages()
	
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

// SetLang sets the current language
func SetLang(lang string) error {
	currentLang = lang
	return loadMessages()
}

// SupportedLanguages returns a list of supported language codes
func SupportedLanguages() []string {
	return []string{"en", "ja"}
}
