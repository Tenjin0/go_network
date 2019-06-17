package chap3

import (
	"fmt"
	"net"
)

func ThreadedIPEchoServer() {

	service := ":1200"
	listener, err := net.Listen("tcp", service)
	checkError(err, 1)
	fmt.Println("Listen on port", 1200)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
