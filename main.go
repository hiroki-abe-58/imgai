package main

import (
	"fmt"
	"os"
)

const version = "0.1.0"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Printf("imgai version %s\n", version)
		return
	}

	fmt.Println("🎨 imgai - AI-powered image processing CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  imgai --version    Show version information")
	fmt.Println()
	fmt.Println("More features coming soon!")
}
