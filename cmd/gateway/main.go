package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: jai-router <command>")
		fmt.Println("commands: start, init, stats")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "start":
		fmt.Println("starting jai-router gateway...")
		// TODO: wire up HTTP server
	case "init":
		fmt.Println("initializing jai-router config...")
		// TODO: first-run setup
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
