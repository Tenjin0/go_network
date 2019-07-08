package main

import (
	"bufio"
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
	conn.Write([]byte(PWD))

	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	fmt.Println("Current dir \"" + s + "\"")
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

		// lise trailing whitespace
		line = strings.TrimRight(line, " \t\r\n")
		strs := strings.SplitN(line, " ", 2)

		switch strs[0] {
		case uiPwd:
			pwdRequest(conn)

		case uiQuit:
			conn.Close()
			os.Exit(0)
		}
	}
}
