package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func ProxyGet(args []string) {

	if len(args) < 2 {
		fmt.Println("Usage: make run target=ProxyGet host=http://host:port proxy=http://host:port")

		os.Exit(1)
	}

	proxyString := args[1]

	proxyURL, err := url.Parse(proxyString)
	checkError(err)

	rawUrl := args[0]

	url, err := url.Parse(rawUrl)
	checkError(err)

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	// client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", url.String(), nil)

	checkError(err)

	urlp, _ := transport.Proxy(request)
	fmt.Println("Proxy", urlp)

	dump, _ := httputil.DumpRequest(request, false)

	fmt.Println(string(dump))
}
