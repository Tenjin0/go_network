package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func ReadHtml(file string) {

	bytes, err := ioutil.ReadFile("html/" + file)
	checkError(err)

	r := strings.NewReader(string(bytes))

	z := html.NewTokenizer(r)

	depth := 0

	for {

		tt := z.Next()

		for n := 0; n < depth; n++ {
			fmt.Print(" ")
		}

		switch tt {
		case html.ErrorToken:
			fmt.Println("Error ", z.Err().Error())
			os.Exit(0)
		case html.TextToken:
			fmt.Println("Text: \"" + strings.Trim(z.Token().String(), " ") + "\"")
		case html.StartTagToken, html.EndTagToken:

			fmt.Println("Tag: \"" + z.Token().String() + "\"")
			if tt == html.StartTagToken {
				depth++
			} else {
				depth--
			}
		}
	}
}
