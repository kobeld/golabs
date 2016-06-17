package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Http server with ListenAndServer call")
	})

	println("Start port on :6001")
	err := http.ListenAndServe(":6001", nil)
	if err != nil {
		panic(err)
	}

}
