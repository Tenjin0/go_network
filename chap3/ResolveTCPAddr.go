package main
import (
	"fmt"
	"net"
	"os"
)

func ResolveTCPAddr(args []string) {

	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s net host port", "ResolveTCPAddr")
		os.Exit(1)
	}

	network := args[0]
	host := args[1]
	port := args[2]

	tcpAddr, err := net.ResolveTCPAddr(network, host+":"+port)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(2)
	}

	fmt.Println(tcpAddr)
}
