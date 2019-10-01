package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func JSONArithServer() {

	arith := new(Arith)
	rpc.Register(arith)
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")

	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}

}
