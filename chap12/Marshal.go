package main

import (
	"encoding/xml"
	"fmt"
)

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string "attr"
	Address string "chardata"
}

type Person struct {
	Name  Name
	Email []Email
}

func Marshal() {

	person := Person{
		Name: Name{
			Family:   "Petit",
			Personal: "Patrice",
		},
		Email: []Email{
			Email{
				Kind:    "home",
				Address: "petitpatrice@hotmail.com",
			},
			Email{
				Kind:    "work",
				Address: "petitpatrice@gmail.com",
			},
		},
	}

	buff, _ := xml.Marshal(person)
	fmt.Println(string(buff))

}
