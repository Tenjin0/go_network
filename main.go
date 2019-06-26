package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Tenjin0/go_network/chap3"
	"github.com/Tenjin0/go_network/chap4"
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
	case "GetHeadInfo":
		chap3.GetHeadInfo(args)
	case "DayTimeServer":
		chap3.DayTimeServer()
	case "EchoServer":
		chap3.EchoServer()
	case "UDPDayTimeClient":
		chap3.UDPDayTimeClient(args)
	case "UDPDayTimeServer":
		chap3.UDPDayTimeServer()
	case "IPGetHeadInfo":
		chap3.IPGetHeadInfo(args)
	case "ThreadedIPEchoServer":
		chap3.ThreadedIPEchoServer()
	case "AsnDayTimeServer":
		chap4.DaytimeServer()
	case "AsnDayTimeClient":
		chap4.ASNDaytimeClient(args)
	case "SaveJSON":
		chap4.SaveJSON(args)
	case "LoadJSON":
		chap4.LoadJSON(args)
	case "EchoJSONClient":
		chap4.EchoJSONClient(args)
	case "EchoJSONServer":
		chap4.EchoJSONServer()
	case "LoadGob":
		chap4.LoadGob()
	case "SaveGob":
		chap4.SaveGob()
	case "EchoGobClient":
		chap4.EchoGobClient(args)
	case "EchoGobServer":
		chap4.EchoGobServer()
	case "Base64":
		chap4.Base64()
	default:
		fmt.Fprintln(os.Stderr, "No", name, "available")
	}

}
