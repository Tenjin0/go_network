package chap3

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func checkError(err error, i int) {

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(i)
	}
}

func GetHeadInfo(args []string) {

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: %s host:port", "GetHeadInfo")
		os.Exit(1)
	}

	service := args[0]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err, 1)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, 2)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err, 3)

	result, err := ioutil.ReadAll(conn)
	checkError(err, 4)

	fmt.Println(string(result))
	os.Exit(0)
}
