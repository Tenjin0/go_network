package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s \"function\" ...args\n", filepath.Base(os.Args[0]))

		os.Exit(1)
	} else {
		name := os.Args[1]
		// args := os.Args[2:]

		switch name {
		case "FTPServer":
			FTPServer()
		case "FTPClient":
			FTPClient()
		case "Encoding":
			Encoding()
		case "UTF16Encoding":
			UTF16Encoding()
		default:
			fmt.Fprintln(os.Stderr, "No", name, "available")
		}
	}
}
