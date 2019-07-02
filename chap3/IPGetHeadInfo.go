package main
import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [100]byte
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

func IPGetHeadInfo(args []string) {

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", "IPGetHeadInfo")
		os.Exit(1)
	}

	service := args[0]

	conn, err := net.Dial("tcp", service)
	checkError(err, 2)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err, 3)

	result, err := readFully(conn)
	checkError(err, 4)

	fmt.Println(string(result))

	os.Exit(0)
}
