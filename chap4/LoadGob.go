package main
import (
	"encoding/gob"
	"fmt"
	"os"
)

func loadGob(filename string, key interface{}) {

	inFile, err := os.Open(filename)
	checkError(err, 1)

	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err, 2)

	inFile.Close()
}

func LoadGob() {

	var person Person

	loadGob("person.gob", &person)

	fmt.Println("Person", person.String())
}
