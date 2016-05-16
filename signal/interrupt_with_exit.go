package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func handleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for _ = range c {
		log.Println("Receive interrupt signal")
		os.Exit(0)
	}
}

func main() {

	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Pong")
	})

	go handleInterrupt()

	log.Println("Listen and serve on port: 9090")
	http.ListenAndServe(":9090", nil)
}
