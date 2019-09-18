package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func acceptableCharset(contentTypes []string) bool {

	for _, cType := range contentTypes {
		if strings.Index(cType, "utf-8") != 1 {
			return true
		}
	}
	return false
}

func Get(args []string) {

	if len(args) < 1 {
		fmt.Println("Usage: make run target=Get host=http://host:port")
		os.Exit(1)
	}

	url := args[0]

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
	}

	if response.StatusCode != 200 {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	fmt.Println("The response header is")

	b, _ := httputil.DumpResponse(response, false)

	fmt.Print(string(b))

	contentTypes := response.Header["Content-Type"]

	if !acceptableCharset(contentTypes) {
		fmt.Println("Cannnot handle", contentTypes)
		os.Exit(4)
	}

	fmt.Println("The response body is")

	var buf [512]byte
	reader := response.Body

	for {
		n, err := reader.Read(buf[0:])
		fmt.Print(string(buf[0:n]))
		if err != nil {
			break
		}
	}
}
