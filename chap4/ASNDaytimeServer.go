package main
import (
	"encoding/asn1"
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error, i int) {

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(i)
	}
}

func handleClientAsn(conn net.Conn) {
	defer conn.Close()

	daytime := time.Now()

	mdata, _ := asn1.Marshal(daytime)
	conn.Write(mdata)
}

func DaytimeServer() {

	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err, 1)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, 2)

	fmt.Fprintf(os.Stderr, "Listen on %s", tcpAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClientAsn(conn)
	}

}
