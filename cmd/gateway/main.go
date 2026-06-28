package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jayesh1126/jai-router/internal/proxy"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: jai-router <command>")
		fmt.Println("commands: start, init, stats")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "start":
		startServer()
		// TODO: wire up HTTP server
	case "init":
		fmt.Println("initializing jai-router config...")
		// TODO: first-run setup
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func startServer() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY environment variable is required")
	}

	baseURL := os.Getenv("OPENAI_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.openai.com/v1"
	}

	upstream, err := proxy.NewProxy(baseURL, apiKey)
	if err != nil {
		log.Fatalf("invalid upstream base URL: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/v1/chat/completions", upstream)
	mux.Handle("/v1/models", upstream)
	mux.Handle("/v1/embeddings", upstream)

	addr := ":8080"
	fmt.Println("jai-router listening on http://localhost" + addr)
	fmt.Println("  OpenAI endpoint: http://localhost" + addr + "/v1")
	fmt.Println("  forwarding to:   " + baseURL)
	log.Fatal(http.ListenAndServe(addr, mux))
}
