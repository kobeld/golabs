package main

import "net"

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	buf := make([]byte, 1024)

	n, err := c.Read(buf)
	if err != nil {
		panic(err)
	}

	println("Received message:")
	println(n)
	println(string(buf))
}
