package chap4

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

func EchoGobClient(args []string) {

	if len(args) != 1 {
		fmt.Println("Usage: go run main.go EchoGobServer host:port")
		os.Exit(1)
	}

	service := args[0]

	conn, err := net.Dial("tcp", service)
	checkError(err, 2)

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	person := GetPerson()
	encoder.Encode(person)

	var newPerson Person
	decoder.Decode(&newPerson)

	fmt.Println(newPerson.String())

	os.Exit(0)
}
