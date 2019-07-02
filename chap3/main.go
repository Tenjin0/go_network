package main

import (
	"fmt"
	"os"
	"path/filepath"
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
		IP(args)
	case "Mask":
		Mask(args)
	case "ResolveIP":
		ResolveIP(args)
	case "LookUpHost":
		LookUpHost(args)
	case "LookUpPort":
		LookUpPort(args)
	case "ResolveTCPAddr":
		ResolveTCPAddr(args)
	case "GetHeadInfo":
		GetHeadInfo(args)
	case "DayTimeServer":
		DayTimeServer()
	case "EchoServer":
		EchoServer()
	case "UDPDayTimeClient":
		UDPDayTimeClient(args)
	case "UDPDayTimeServer":
		UDPDayTimeServer()
	case "IPGetHeadInfo":
		IPGetHeadInfo(args)
	case "ThreadedIPEchoServer":
		ThreadedIPEchoServer()
	default:
		fmt.Fprintln(os.Stderr, "No", name, "available")
	}

}
