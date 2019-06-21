package chap4

import (
	"encoding/json"
	"fmt"
	"os"
)

func (p Person) String() string {

	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func loadJSON(filename string, key interface{}) {

	inFile, err := os.Open(filename)
	checkError(err, 2)

	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err, 3)

	inFile.Close()
}

func LoadJSON(args []string) {

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: go run main.go LoadJSON \"filename\"")
		os.Exit(1)
	}

	filename := args[0]

	var person Person
	loadJSON(filename, &person)

	fmt.Println("Person", person.String())
}
