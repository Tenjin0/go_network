package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func Get(args []string) {

	if len(args) < 1 {
		fmt.Println("Usage: make run target=Get host=http://host:port")
		os.Exit(1)
	}

	url := args[0]

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	if response.StatusCode != 200 {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	fmt.Println("The response header is")
	b, _ := httputil.DumpResponse(response, false)
	fmt.Print(string(b))

	contentTypes := response.Header["Content-Type"]
}
