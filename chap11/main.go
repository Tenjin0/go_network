package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "file")
		os.Exit(1)
	}

	ReadHtml(os.Args[1])

}
