package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"strconv"
)

func TLSEchoClient(service string) {
	fmt.Println(service)
	certPemFile, err := os.Open("certs/patrice.petit.name.pem")
	checkError(err)

	pemBytes := make([]byte, 1000)

	_, err = certPemFile.Read(pemBytes)
	checkError(err)

	defer certPemFile.Close()

	certPool := x509.NewCertPool()

	ok := certPool.AppendCertsFromPEM(pemBytes)

	if !ok {
		fmt.Println("PEM read failed")
	} else {
		fmt.Println("PEM read ok")
	}

	conn, err := tls.Dial("tcp", service, &tls.Config{RootCAs: certPool})
	checkError(err)

	for n := 0; n < 10; n++ {

		stringToSend := "Hello" + strconv.Itoa(n)

		fmt.Println("Writing...", stringToSend)
		conn.Write([]byte(stringToSend))

		var buf [512]byte

		n, err := conn.Read(buf[0:])
		checkError(err)

		fmt.Println(string(buf[0:n]))
	}
	conn.Close()
	os.Exit(0)
}
