package main
import (
	"encoding/gob"
	"os"
)

func saveGob(filename string, key interface{}) {

	outFile, err := os.Create(filename)
	checkError(err, 1)

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err, 2)

	outFile.Close()
}

func SaveGob() {

	person := Person{
		Name: Name{
			Personal: "Patrice",
			Family:   "Petit",
		},
		Email: []Email{
			Email{
				Kind:    "home",
				Address: "patricepetit@hotmail.com",
			},
			Email{
				Kind:    "work",
				Address: "pp@a.com",
			},
		},
	}

	saveGob("person.gob", person)
}
