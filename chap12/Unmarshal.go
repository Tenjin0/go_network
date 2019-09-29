package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type UnName struct {
	Family   string `xml:"family"`
	Personal string `xml:"personal"`
}

type UnEmail struct {
	Type    string `xml:"type,attr"`
	Address string `xml:",chardata"`
}
type UnPerson struct {
	XMLName UnName    `xml:"person"`
	Name    UnName    `xml:"name"`
	Email   []UnEmail `xml:"email"`
}

func Unmarshal() {

	xmlfile, err := ioutil.ReadFile("xml/person.xml")
	checkError(err)

	var person UnPerson
	err = xml.Unmarshal(xmlfile, &person)
	checkError(err)

	fmt.Println(person.Name)
}
