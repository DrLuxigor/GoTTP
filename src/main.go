package main

import (
	"bufio"
	"fmt"
	"net"

	"lukaskofler.dev/gottp/src/pkg/http"
)

var (
	port           = 8050
	maxRequestSize = 1 << 24
	version        = "0.0.1"
)

func main() {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("GoTTP v%s listening on port %d\n", version, port)

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	request := http.ParseHttpRequest(reader)
	request.Print(false)

	conn.Write([]byte("HTTP/1.1 200 OK\nServer: GoTTP\nContent-Type: text/html\nContent-Lenght: 21\n\n<h1>Hello World</h1>\n"))
}
