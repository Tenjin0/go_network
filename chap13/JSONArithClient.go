package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"
)

func JSONArithClient(args []string) {

	if len(args) != 2 {
		fmt.Println("Usage: go run . JSONClient server:port")
		os.Exit(1)
	}

	service := args[1]
	client, err := jsonrpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	params := Values{17, 8}
	var reply int

	err = client.Call("Arith.Multiply", params, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", params.A, params.B, reply)

	var quot Quotient

	err = client.Call("Arith.Divide", params, &quot)
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", params.A, params.B, quot.Quo, quot.Rem)
}
