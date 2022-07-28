package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SVR_HOST = "localhost"
	SVR_PORT = "9982"
	SVR_TYPE = "tcp"
)

func main() {
	fmt.Println("server is running")
	svr, err := net.Listen(SVR_TYPE, SVR_HOST+":"+SVR_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer svr.Close()
	fmt.Println("Listening on: " + SVR_HOST + ":" + SVR_PORT)
	fmt.Println("Waiting for connection...")

	for {
		conn, err := svr.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			os.Exit(1)
		}
		if n == 0 {
			break
		}
		fmt.Print(string(buf[0:]))
	}
}
