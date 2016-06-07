package main

import (
	"net"
	"os"
	"time"
)

func read(conn net.Conn) {

	buf := make([]byte, 512)
	for {

		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}

		println("Client got: ", string(buf[0:n]))
	}
}

func main() {

	var (
		sockPath = os.Getenv("GOPATH") + "/src/github.com/kobeld/golabs/sockets/echo.sock"
	)

	conn, err := net.Dial("unix", sockPath)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go read(conn)

	for {
		_, err := conn.Write([]byte("Hello Unix Domain Socket"))
		if err != nil {
			panic(err)
		}

		time.Sleep(3 * time.Second)
	}
}
