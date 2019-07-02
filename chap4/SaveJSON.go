package main
import (
	"encoding/json"
	"fmt"
	"os"
)

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

type Person struct {
	Name  Name
	Email []Email
}

func saveJSON(filename string, key interface{}) {

	outFile, err := os.Create(filename)
	checkError(err, 1)

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(key); err != nil {
		checkError(err, 2)
	}
	outFile.Close()
}

func SaveJSON(args []string) {

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s filename", "SaveJSON")
	}

	person := Person{
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
				Address: "petitpatrice@gmail.com",
			},
		},
	}

	fmt.Println(person)

	filname := args[0]
	saveJSON(filname, person)
}
