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

	rawURL := args[0]

	url, err := url.Parse(rawURL)
	checkError(err)

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}

	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", url.String(), nil)

	checkError(err)

	urlp, _ := transport.Proxy(request)
	fmt.Println("Proxy", urlp)

	dump, _ := httputil.DumpRequest(request, false)

	fmt.Println(string(dump))

	response, err := client.Do(request)

	if response.StatusCode != 200 {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	fmt.Println(response.Status)

	var buf [512]byte
	reader := response.Body

	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}

	os.Exit(0)
}
