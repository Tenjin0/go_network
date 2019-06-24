package chap4

import (
	"encoding/json"
	"fmt"
	"net"
)

func EchoServer(args []string) {

	service := "0.0.0.0:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err, 1)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, 2)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		encoder := json.NewEncoder(conn)
		decoder := json.NewDecoder(conn)

		for n := 0; n < 10; n++ {
			var person Person
			decoder.Decode(&person)
			fmt.Println(person.String())
			encoder.Encode(person)
		}

		conn.Close()
	}

}
