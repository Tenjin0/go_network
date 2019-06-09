package chap3

import (
	"fmt"
	"net"
	"os"
)

func IP(args []string) {

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s \"ip-addr\"\n", "IP")
		os.Exit(1)
	}
	name := args[0]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", addr)
	}
	os.Exit(0)
}
