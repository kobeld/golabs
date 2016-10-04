package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	n, err := fmt.Fprintf(conn, "Hello Aaron!")
	if err != nil {
		panic(err)
	}

	println("Sent")
	println(n)

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	println(string(buf))

}
