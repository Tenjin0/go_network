package chap4

import (
	"fmt"

	"github.com/Tenjin0/go_network/chap4/src/person"
	"github.com/golang/protobuf/proto"
)

func ProtocolBuffer() {

	name := person.Person_Name{
		Family:   "Petit",
		Personal: "Patrice",
	}

	email1 := person.Person_Email{
		Kind:    "home",
		Address: "patricepetit@hotmail.com",
	}
	email2 := person.Person_Email{
		Kind:    "work",
		Address: "p@p@a.com",
	}

	emails := []*person.Person_Email{&email1, &email2}

	p := person.Person{
		Name:  &name,
		Email: emails,
	}

	data, err := proto.Marshal(&p)
	checkError(err, 1)

	newP := person.Person{}

	err = proto.Unmarshal(data, &newP)
	checkError(err, 2)

	fmt.Println(newP)
}
