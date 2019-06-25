package chap4

import (
	"encoding/json"
	"fmt"
	"net"
)

func handleClientJson(conn net.Conn) {

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	var person Person
	decoder.Decode(&person)
	fmt.Println(person.String())
	encoder.Encode(person)

	conn.Close()
}

func EchoJSONServer() {

	service := "0.0.0.0:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err, 1)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, 2)
	fmt.Println("Listening to", tcpAddr.String())

	for {
		conn, err := listener.Accept()
		fmt.Printf("New connection from %s\n\n", conn.RemoteAddr())
		if err != nil {
			continue
		}
		go handleClientJson(conn)

	}

}
