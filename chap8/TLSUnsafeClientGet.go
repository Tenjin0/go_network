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
}
