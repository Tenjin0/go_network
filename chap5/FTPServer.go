package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
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
	result := bytes.NewBuffer(nil)

	fmt.Println("handleClient")

	for {
		n, err := conn.Read(buf[0:])

		if err != nil {
			fmt.Println("find an error", err.Error())
			conn.Close()
			return
		}

		result.Write(buf[0:n])

		contents := string(result.Bytes())

		i := strings.Index(contents, "\n")

		if i >= 0 {
			result.Reset()
			if len(contents) >= 2 && strings.ToUpper(contents[0:2]) == CD {
				fmt.Println(contents, contents[3:])
				chdir(conn, contents[3:])
			} else if len(contents) >= 3 && strings.ToUpper(contents[0:3]) == DIR {
				dirList(conn)
			} else if len(contents) >= 3 && strings.ToUpper(contents[0:3]) == PWD {
				pwd(conn)
			}
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
	tmp := strings.Trim(s, " \n\r\t")

	if os.Chdir(tmp) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func dirList(conn net.Conn) {

	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	names, err := dir.Readdirnames(-1)

	if err != nil {
		return
	}

	for _, nm := range names {
		conn.Write([]byte(nm + "\r\n"))
	}
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
