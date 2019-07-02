package main
import (
	"fmt"
	"net"
	"os"
)

func LookUpPort(args []string) {

	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s network-type service\n", "LookUpPort")
		os.Exit(1)
	}

	netWorkType := args[0]
	service := args[1]
	port, err := net.LookupPort(netWorkType, service)

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(2)
	}

	fmt.Println("Service port", port)
	os.Exit(0)
}
