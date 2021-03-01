package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//net = network
	nl, err := net.Listen("tcp", ":8888") // ip:port  port range : 0 to 65535
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) // 1 = stop with error
	}

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

	fmt.Println(n)
	reqstr := string(bs) // conversion
	fmt.Println(reqstr)

	body := `<!DOCTYPE html>
	<html>
	<head>
	<title>Page Title</title>
	</head>
	<body>
	
	<h1>Welcome to coding boot-camp</h1>
	<p>by Master Academy</p>
	
	</body>
	</html>`

	resp := "HTTP/1.1 200 OK\r\n" +
		"Content-Length: %d\r\n" +
		"Content-Type: text/html\r\n" +
		"\r\n%s"

	msg := fmt.Sprintf(resp, len(body), body)
	fmt.Println(msg)
	conn.Write([]byte(msg))

	/*
		remoteAddr := conn.RemoteAddr().String()
		fmt.Println(remoteAddr)

		//byte
		conn.Write([]byte("welcome we have received your message"))

		conn.Close()
		nl.Close()
	*/
}
