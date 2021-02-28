package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//net = network
	nl, err := net.Listen("tcp", ":8888") //ip:port
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	remoteAddr := conn.RemoteAddr().String()
	fmt.Println(remoteAddr)

	//byte
	conn.Write([]byte("welcome we have received your message"))

	conn.Close()
	nl.Close()
}
