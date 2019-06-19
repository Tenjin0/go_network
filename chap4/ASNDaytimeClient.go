package chap4

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func readFully(conn net.Conn) ([]byte, error) {

	defer conn.Close()

	var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {

		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}

func ASNDaytimeClient(args []string) {

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", "ASNDaytimeClient")
	}

	service := args[0]

	conn, err := net.Dial("tcp", service)
	checkError(err, 1)

	result, err := readFully(conn)
	checkError(err, 2)

	fmt.Println(result)
}
