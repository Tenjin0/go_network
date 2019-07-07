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

func FTPClient() {

	host := "localhost"
	_, err := net.Dial("tcp", host+":1202")
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

		fmt.Println(strs)

		switch strs[0] {
		case uiPwd:
			pwdRequest(conn)
		}
	}
}
