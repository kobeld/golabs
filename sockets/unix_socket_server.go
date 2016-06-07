package main

import (
	"net"
	"os"
)

func serveConn(c net.Conn) {

	buf := make([]byte, 512)

	for {
		n, err := c.Read(buf)
		if err != nil {
			panic(err)
		}

		data := buf[0:n]
		println("Server got: ", string(data))

		_, err = c.Write(data)
		if err != nil {
			panic(err)
		}
	}
}

func main() {

	var (
		sockFile = os.Getenv("GOPATH") + "/src/github.com/kobeld/golabs/sockets/echo.sock"
	)

	listener, err := net.Listen("unix", sockFile)
	if err != nil {
		panic(err)
	}

	for {
		unixConn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go serveConn(unixConn)
	}
}
