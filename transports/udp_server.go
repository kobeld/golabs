package main

import "net"

func main() {

	ln, err := net.ListenPacket("udp", ":7070")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	println("UDP server started...")

	for {
		buffer := make([]byte, 1024)

		n, addr, err := ln.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}

		println("Received message from:")
		println(n)
		println(addr.String())
		println(string(buffer))

		// time.Sleep(20 * time.Second)
	}
}
