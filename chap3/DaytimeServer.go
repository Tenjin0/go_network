package main

import (
	"fmt"
	"net"
	"time"
)

func DayTimeServer() {

	service := ":1200"

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)

	checkError(err, 1)

	listener, err := net.ListenTCP("tcp", tcpAddr)

	for {
		fmt.Println("wait accepting")
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		daytime := time.Now().String()

		conn.Write([]byte(daytime))
		conn.Close()
	}

}
