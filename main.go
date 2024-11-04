package main

import "fmt"
import "net"

var (
	port = 8050;
	maxRequestSize = 1<<24;
	version = "0.0.1"
)



func main(){
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port));

	if err != nil {
		fmt.Println(err);
		return;
	}
	
	fmt.Println(fmt.Sprintf("GoTTP v%s listening on port %d", version, port));

	for {
		conn, err := ln.Accept();

		if err != nil {
			fmt.Println(err);
			continue;
		}

		go handleConnection(conn);
	}
}



func handleConnection(conn net.Conn) {
	defer conn.Close();

	buf := make([]byte, maxRequestSize);
	_, err := conn.Read(buf);

	if err != nil {
		fmt.Println(err);
		return;
	}

	conn.Write([]byte("HTTP/1.1 200 OK\nServer: GoTTP\nContent-Type: text/html\nContent-Lenght: 6\n\n{Hi!}\n"));
}
