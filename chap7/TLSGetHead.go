package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"os"
)

func TLSGetHead(url string) {

	service := url

	conn, err := tls.Dial("tcp", service, nil)
	checkError(err)

	defer conn.Close()

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}
