package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			fmt.Fprintf(w, "Err: %s", err.Error())
			return
		}

		fmt.Fprintf(w, "From host: %s", hostname)
	})

	println("Start port on :6666")
	err := http.ListenAndServe(":6666", nil)
	if err != nil {
		panic(err)
	}

}
