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
	case "AsnDayTimeServer":
		DaytimeServer()
	case "AsnDayTimeClient":
		ASNDaytimeClient(args)
	case "SaveJSON":
		SaveJSON(args)
	case "LoadJSON":
		LoadJSON(args)
	case "EchoJSONClient":
		EchoJSONClient(args)
	case "EchoJSONServer":
		EchoJSONServer()
	case "LoadGob":
		LoadGob()
	case "SaveGob":
		SaveGob()
	case "EchoGobClient":
		EchoGobClient(args)
	case "EchoGobServer":
		EchoGobServer()
	case "Base64":
		Base64()
	case "ProtocolBuffer":
		ProtocolBuffer()
	default:
		fmt.Fprintln(os.Stderr, "No", name, "available")
	}

}
