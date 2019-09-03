package main

import (
	"fmt"
	"net/http"
	"os"
)

func Head(args []string) {

	if len(args) < 1 {
		fmt.Println("Usage: make run target=Head host=http://host:port")
		os.Exit(1)
	}

	url := args[0]

	response, err := http.Head(url)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println(response.Request)

	for k, v := range response.Header {
		fmt.Println(k+":", v)
	}
	os.Exit(0)
}
