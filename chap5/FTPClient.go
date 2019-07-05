package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func FTPClient(args []string) {

	if len(args) != 2 {
		fmt.Println("Usage: ", args[0], "host")
		os.Exit(1)
	}

	host := args[0]
	_, err := net.Dial("tcp", host+":1202")
	checkError(err)

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		fmt.Println(line)
		if err != nil {
			break
		}

		// lise trailing whitespace
		line = strings.TrimRight(line, " \t\r\n")
		fmt.Println(line)

	}
}
