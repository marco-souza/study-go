package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/zserge/webview"
)

func main() {
	// Define port
	port := ":8071"

	// Start server
	startServer(port)

	// Start window
	startClient(port)
}

func startServer(port string) {
	http.HandleFunc("/", rootHandler)

	// start server
	go func() {
		log.Fatal(http.ListenAndServe(port, nil))
	}()
}

func startClient(port string) {
	// Set variables
	title := "Minimal webview example"
	width, height := size("fullhd")
	resizable := true
	ip := "127.0.0.1"
	url := "http://" + ip + port

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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello motherfucker, %q", html.EscapeString(r.URL.Path))
}
