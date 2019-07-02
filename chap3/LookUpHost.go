package main
import (
	"fmt"
	"net"
	"os"
)

func LookUpHost(args []string) {

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", "LookUpHost")
		os.Exit(1)
	}

	name := args[0]
	addrs, err := net.LookupHost(name)

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(2)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}

	os.Exit(0)
}
