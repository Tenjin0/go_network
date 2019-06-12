package chap3

import (
	"fmt"
	"net"
	"time"
)

func handleClientUDP(conn *net.UDPConn) {

	var buf [512]byte
	fmt.Println("Read")
	_, addr, err := conn.ReadFromUDP(buf[0:])

	if err != nil {
		return
	}

	daytime := time.Now().String()

	conn.WriteToUDP([]byte(daytime), addr)
}

func UDPDayTimeServer() {

	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err, 1)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err, 2)

	for {
		handleClientUDP(conn)
	}
}
