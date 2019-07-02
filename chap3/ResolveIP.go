package main
import (
	"fmt"
	"net"
	"os"
)

func ResolveIP(args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", "ResolveIP")
		os.Exit(1)
	}

	name := args[0]
	addr, err := net.ResolveIPAddr("ip", name)

	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(0)
	}

	fmt.Println("Resolved address is", addr.String())
	os.Exit(0)
}
