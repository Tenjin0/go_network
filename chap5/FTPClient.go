package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	uiDir  = "dir"
	uiCd   = "cd"
	uiPwd  = "pwd"
	uiQuit = "quit"
)

func pwdRequest(conn net.Conn) {

	conn.Write([]byte(PWD + "\n"))

	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	fmt.Println("Current dir \"" + s + "\"")
}

func dirRequest(conn net.Conn) {

	conn.Write([]byte(DIR + "\n"))

	var buf [512]byte

	result := bytes.NewBuffer(nil)

	for {
		n, _ := conn.Read(buf[:])
		result.Write(buf[0:n])
		length := result.Len()
		contents := result.Bytes()
		if string(contents[length-4:]) == "\r\n\r\n" {
			fmt.Println(string(contents[0 : length-4]))
			return
		}
	}

}

func cdRequest(conn net.Conn, path string) {

	conn.Write([]byte(CD + " " + path + "\n"))

	var response [512]byte
	n, _ := conn.Read(response[0:])

	s := string(response[0:n])

	fmt.Println("change dir", s)
}

func FTPClient() {

	host := "localhost"
	conn, err := net.Dial("tcp", host+":1202")
	checkError(err)

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimRight(line, " \t")
		strs := strings.SplitN(line, " ", 2)

		switch strings.TrimRight(strs[0], " \t\r\n") {
		case uiPwd:
			pwdRequest(conn)
		case uiDir:
			dirRequest(conn)
		case uiCd:
			cdRequest(conn, strs[1])
		case uiQuit:
			conn.Close()
			os.Exit(0)
		}
	}
}
