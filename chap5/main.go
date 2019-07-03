package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s \"function\" ...args", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	name := os.Args[1]
	// args := os.Args[2:]
	switch name {
	case "FTPServer":
		FTPServer()
	default:
		fmt.Fprintln(os.Stderr, "No", name, "available")
	}

}
