package main

import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
)

const BOM = '\ufffe'

func readShorts(conn net.Conn) ([]uint16, error) {

	var buf [512]byte

	_, err := conn.Read(buf[0:2])

	if err != nil {
		return make([]uint16, 0), err // []uint16{}
	}

	return []uint16{}, nil

}

func UTF16Client(args []string) {

	if len(args) != 1 {
		fmt.Println("Usage: make run target=UTF16Client \"host:port\"")
		os.Exit(1)
	}

	service := args[0]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	shorts, err := readShorts(conn)

	if err != nil {
		fmt.Println(err.Error())
	}
	ints := utf16.Decode(shorts)
	str := string(ints)

	fmt.Println(str)

	os.Exit(0)

}
