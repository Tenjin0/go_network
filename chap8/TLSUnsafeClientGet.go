package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func TLSUnsafeClientGet(args []string) {

	if len(args) != 1 {
		fmt.Println("Usage: make run target=TLSUnsafeClientGet host=https://host:port/page")
		os.Exit(1)
	}

	url, err := url.Parse(args[0])
	checkError(err)

	if url.Scheme != "https" {
		fmt.Println("No https scheme ", url.Scheme)
		os.Exit(1)
	}

	transport := &http.Transport{}
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)

	response, err := client.Do(request)
	checkError(err)

	if response.StatusCode != 200 {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	fmt.Println("get a response")

	chSet := getCharset(response)
	fmt.Printf("got charset %S\n", chSet)

	if chSet != "UTF-8" {
		fmt.Printf("Cannot handle", chSet)
		os.Exit(4)
	}

	var buf [512]byte

	reader := response.Body

	fmt.Println("got body")

	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}

}
