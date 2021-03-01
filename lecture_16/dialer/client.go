package main

import (
	"fmt"
	"net"
)

func main() {

	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", ":8888") //remote server ip:91.205.173.170

	if err != nil {
		fmt.Println(err.Error())
	}

	// way-1
	//bs := []byte("Hello")
	//conn.Write(bs)

	// way-2
	conn.Write([]byte("Hello"))

	bs := make([]byte, 1024)
	n, err := conn.Read(bs)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(bs[:n]))

	conn.Close()
}
