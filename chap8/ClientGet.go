package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func ClientGet(args []string) {

	if len(args) != 2 {
		fmt.Println("Usage: make run target=ClientGet host=http://host:port/page")
	}

	url, err := url.Parse(args[0])
	checkError(err)

	client := &http.Client{}

	request, err := http.NewRequest("HEAD", url.String(), nil)
	checkError(err)
	request.Header.Add("Accept-Charset", "uft-8;q=1, ISO-8859-1;q=0")

	response, err := client.Do(request)
	checkError(err)

	if response.StatusCode != 200 {
		fmt.Println(response.Status)
		os.Exit(2)
	}
}
