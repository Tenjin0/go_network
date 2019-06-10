package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Tenjin0/go_network/chap3"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s \"function\" ...args", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	name := os.Args[1]
	args := os.Args[2:]
	switch name {
	case "IP":
		chap3.IP(args)
	case "Mask":
		chap3.Mask(args)
	case "ResolveIP":
		chap3.ResolveIP(args)
	case "LookUpHost":
		chap3.LookUpHost(args)
	case "LookUpPort":
		chap3.LookUpPort(args)
	case "ResolveTCPAddr":
		chap3.ResolveTCPAddr(args)
	default:
		fmt.Fprintln(os.Stderr, "No", name, "available")
	}

}
