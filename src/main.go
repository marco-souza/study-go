package main

import (
	"fmt"

	"github.com/zserge/webview"
)

func main() {
	title := "Minimal webview example"
	width, height := size("fullhd")
	resizable := true
	url := "https://google.com"

	fmt.Println(height, width)

	// Open wikipedia in a 800x600 resizable window
	webview.Open(title, url, width, height, resizable)
}

func size(format string) (int, int) {
	switch format {
	case "fullhd":
		return 1920, 1080

	case "widescreen":
		return 1280, 800

	default:
		return 800, 600
	}
}
