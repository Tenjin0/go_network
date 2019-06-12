package chap3

import (
	"fmt"
	"net"
	"os"
)

func UDPDayTimeClient(args []string) {

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", args[0])
	}

	service := args[0]

	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err, 1)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err, 2)

	_, err = conn.Write([]byte("anything"))
	checkError(err, 3)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err, 4)

	fmt.Println(string(buf[0:n]))

	os.Exit(0)

}
