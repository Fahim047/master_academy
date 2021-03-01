package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	//net = network
	nl, err := net.Listen("tcp", ":8888") // ip:port  port range : 0 to 65535
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) // 1 = stop with error
	}

	defer nl.Close() //await
	log.Printf("server started on :8888")

	for {

		conn, err := nl.Accept()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		bs := make([]byte, 1024)

		n, err := conn.Read(bs)
		if err != nil {
			fmt.Println(err.Error())
		}

		//fmt.Println(n)
		//fmt.Println(bs)
		reqstr := string(bs[:n]) // conversion
		fmt.Println(reqstr)
		recvTime := time.Now().Format("2006-01-02 15:04:05")

		msg := fmt.Sprintf(`Your message: %s, received at %s`, reqstr, recvTime)
		fmt.Println(msg)
		conn.Write([]byte(msg))
	}

}
