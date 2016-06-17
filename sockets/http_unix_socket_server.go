package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

var (
	sockFile = os.Getenv("GOPATH") + "/src/github.com/kobeld/golabs/sockets/http.sock"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	// This here is an example of Chuck Han https://groups.google.com/forum/#!topic/golang-nuts/41TPj4PWBI8
	// Listen on unix socket /tmp/ango.sock

	l, err := net.Listen("unix", sockFile)
	if err != nil {
		panic(err)
	}

	err = http.Serve(l, nil)
	if err != nil {
		panic(err)
	}
}
