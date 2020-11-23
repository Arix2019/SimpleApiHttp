// SERVER
// implementando http manualmente.
package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Aguardando conexões...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go response(conn)
	}
}
