package main

import (
	"fmt"
	"net"
	"os"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error:", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {

	defer conn.Close()

	var buf [512]byte

	for {

		n, err := conn.Read(buf[0:])
		if err != nil {
			conn.Close()
			return
		}

		s := string(buf[0:n])

		if len(s) >= 2 && s[0:2] == CD {
			chdir(conn, s[3:])
		} else if len(s) >= 3 && s[0:3] == DIR {
			dirList(conn)
		} else if len(s) >= 3 && s[0:3] == PWD {
			pwd(conn)
		}
	}
}

func pwd(conn net.Conn) {

	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(s))
}

func chdir(conn net.Conn, s string) {

	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func dirList(conn net.Conn) {

	// dir, err := os.Open(".")
	// if err != nil {

	// }
}

func FTPServer() {

	service := "0.0.0.0:1202"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)

	fmt.Println("Listen on", tcpAddr)

	for {
		conn, err := listener.Accept()

		fmt.Println(conn.LocalAddr())
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}
