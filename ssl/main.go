package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, 鲑鱼~")
	})

	err := http.ListenAndServeTLS(":5555", "iguiyu.com.crt", "iguiyu.com.key", nil)
	if err != nil {
		panic("Http ListenAndServeTLS: " + err.Error())
	}
}
