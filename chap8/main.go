package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}

func main() {

	if len(os.Args) < 2 {

		fmt.Fprintf(os.Stderr, "Usage: %s \"function\" ... args\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	} else {

		name := os.Args[1]
		args := os.Args[2:]
		switch name {
		case "Head":
			Head(args)
		case "Get":
			Get(args)
		case "ClientGet":
			ClientGet(args)
		case "ProxyGet":
			ProxyGet(args)
		case "FileServer":
			FileServer()
			return
		default:
			fmt.Fprintln(os.Stderr, "No", name, "available")
		}
	}
}
