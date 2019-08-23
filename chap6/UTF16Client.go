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

	n, err := conn.Read(buf[0:2])

	if err != nil {
		return make([]uint16, 0), err // []uint16{}
	}

	for {
		m, err := conn.Read(buf[n:])
		if m == 0 || err != nil {
			fmt.Println(m, err)
			break
		}
		n += m
	}

	checkError(err)

	fmt.Println(buf)

	var shorts []uint16

	shorts = make([]uint16, n/2)

	if buf[0] == 0xff && buf[1] == 0xfe {
		// big endian
		for i := 2; i < n; i += 2 {
			shorts[i/2-1] = uint16(buf[i])<<8 + uint16(buf[i+1])
		}

	} else if buf[0] == 0xfe && buf[1] == 0xff {
		// little endian
		for i := 0; i < n; i += 2 {
			shorts[i/2-1] = uint16(buf[i+1])<<8 + uint16(buf[i])
		}
	} else {
		// unknow byte order
		return []uint16{}, new Error
	}

	fmt.Println(shorts)

	return shorts, nil

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
