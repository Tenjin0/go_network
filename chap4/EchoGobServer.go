package main
import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

func EchoGobServer() {

	service := "0.0.0.0:1200"

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err, 1)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, 2)
	fmt.Println("Listening to", tcpAddr.String())

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}
		fmt.Printf("New connection from %s\n\n", conn.RemoteAddr())

		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

		var person Person
		decoder.Decode(&person)
		fmt.Println(person.String())
		encoder.Encode(person)

		conn.Close()
	}

	os.Exit(0)
}
