package chap3

import (
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {

	defer conn.Close()
	var buf [2]byte

	for {
		n, err := conn.Read(buf[0:])

		if err != nil {
			return
		}

		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func EchoServer() {

	service := ":1201"

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)

	checkError(err, 1)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, 2)
	fmt.Println("listening to", tcpAddr)

	for {
		fmt.Println("wait accepting")
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		go handleClient(conn)
	}

}
