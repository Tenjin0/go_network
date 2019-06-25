package chap4

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func GetPerson() Person {

	return Person{
		Name: Name{
			Family:   "Petit",
			Personal: "Patrice",
		},
		Email: []Email{
			Email{
				Kind:    "home",
				Address: "patricepetit@hotmail.com",
			},
			Email{
				Kind:    "work",
				Address: "ppetit@a.com",
			},
		},
	}
}

func EchoJSONClient(args []string) {

	if len(args) != 1 {
		fmt.Println("Usage: go run main.go EchoClient host:port")
		os.Exit(1)
	}

	person := GetPerson()

	service := args[0]

	conn, err := net.Dial("tcp", service)
	checkError(err, 1)

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	encoder.Encode(person)
	var newPerson Person

	decoder.Decode(&newPerson)

	fmt.Println(newPerson.String())

	conn.Close()

	os.Exit(0)
}
