package main
import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func Mask(args []string) {

	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr ones bits\n", "Mask")
		os.Exit(1)
	}
	dotAddr := args[0]
	ones, _ := strconv.Atoi(args[1])
	bits, _ := strconv.Atoi(args[2])

	addr := net.ParseIP(dotAddr)

	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	mask := net.CIDRMask(ones, bits)
	network := addr.Mask(mask)

	fmt.Println("Address is ", addr.String(),
		"\nMask length is", bits,
		"\nLeading ones count is", ones,
		"\nMask is (hex)", mask.String(),
		"\nNetwork is", network.String())
	os.Exit(0)
	os.Exit(0)
}
