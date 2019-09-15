package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")

	if contentType == "" {
		return "utf-8"
	}

	idx := strings.Index(contentType, "charset=")
	if idx == -1 {
		return "utf-8"
	}

	chSet := strings.Trim(contentType[idx+8:], " ")
	return strings.ToLower(chSet)

}

func ClientGet(args []string) {

	if len(args) < 1 {
		fmt.Println("Usage: make run target=ClientGet host=http://host:port/page")
		os.Exit(1)
	}

	url, err := url.Parse(args[0])
	checkError(err)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)
	request.Header.Add("Accept-Charset", "uft-8;q=1, ISO-8859-1;q=0")

	response, err := client.Do(request)
	checkError(err)

	if response.StatusCode != 200 {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	fmt.Println("The response header is")
	b, _ := httputil.DumpResponse(response, false)

	fmt.Print(string(b))

	chSet := getCharset(response)

	fmt.Println(chSet)

}
