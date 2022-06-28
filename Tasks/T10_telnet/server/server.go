package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(fmt.Errorf("error when ResolveTCPAddr: %w", err))
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(fmt.Errorf("error when ListenTCP: %w", err))
	}

	conn, err := listener.Accept()

	if err != nil {
		log.Fatal(fmt.Errorf("error when establish connection: %w", err))
	}

	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)
}
