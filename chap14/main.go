package main

import (
	"fmt"
	"net/http"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . port")
		os.Exit(1)
	}

	port := os.Args[1]

	http.HandleFunc("/", HandleFlashCardSet)
	err := http.ListenAndServe(port, nil)

	checkError(err)
}
