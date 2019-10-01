package main

import (
	"fmt"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Usage: go run main.go Server/Client host?")
	}

	switch os.Args[1] {
	case "Server":
		ArithServer()
		break
	case "Client":
		ArithClient(os.Args[1:])
		break
	case "TCPServer":
		TCPArithServer()
		break
	case "TCPClient":
		TCPArithClient(os.Args[1:])
		break
	case "JSONServer":
		JSONArithServer()
		break
	case "JSONClient":
		JSONArithClient(os.Args[1:])
		break
	default:
		log.Fatal("No registered service", os.Args[1])
	}

}
