package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

func ArithClient(args []string) {

	if len(args) != 2 {
		fmt.Println("Usage: go run . Client serverHostIP")
		os.Exit(1)
	}

	serverHost := args[1]

	client, err := rpc.DialHTTP("tcp", serverHost+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

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
