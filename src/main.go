package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"lukaskofler.dev/gottp/src/pkg/http"
)

var (
	port                 = 8050
	maxRequestSize int64 = 1 << 24
	version              = "0.0.1"
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

	limitreader := io.LimitReader(conn, maxRequestSize)
	reader := bufio.NewReader(limitreader)

	request := http.ParseHttpRequest(reader)
	request.Print(false)

	response := http.HttpResponse{
		Version:     request.Version,
		StatusCode:  200,
		Message:     "OK",
		Headers:     make(map[string]string),
		Cookies:     make([]string, 0),
		ContentType: "text/html",
		Body:        []byte("<h1>Hello World</h1>"),
	}

	response.Headers["Server"] = "GoTTP"
	response.SetCookie("testcookie", "testvalue", map[string]string{"path": "/", "max-age": "3600", "same-site": "Lax", "priority": "High", "http-only": "true", "secure": "true"})

	conn.Write([]byte(response.BuildResponseHeader()))
	if response.Body != nil {
		conn.Write(response.Body)
	}
}
