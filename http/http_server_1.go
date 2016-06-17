package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Http server with seperating listen and server calls")
	})

	l, err := net.Listen("tcp", ":6000")
	if err != nil {
		panic(err)
	}

	println("Start port on :6000")
	err = http.Serve(l, nil)
	if err != nil {
		panic(err)
	}
}
