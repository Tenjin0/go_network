package chap3

import "net"

func ThreadedIPEchoServer() {

	service := ":1200"
	listener, err := net.Listen("tcp", service)
	checkError(err, 1)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
