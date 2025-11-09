package main

import (
	"os"

	"github.com/hiroki-abe-58/imgai/cmd"
	"github.com/hiroki-abe-58/imgai/pkg/i18n"
)

func main() {
	// Initialize i18n system
	lang := os.Getenv("IMGAI_LANG")
	i18n.Init(lang)
	
	cmd.Execute()
}
