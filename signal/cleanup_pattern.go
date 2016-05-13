package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {

	var (
		signalChan = make(chan os.Signal, 1)
		doneChan   = make(chan bool)

		cleanupJob = func() {
			println("Cleaning up services and resources")
		}
	)

	signal.Notify(signalChan, os.Interrupt)

	go func() {

		for _ = range signalChan {
			fmt.Println("\nReceived an interrupt, stopping services...")
			cleanupJob()
			doneChan <- true
		}

	}()

	println("Program started, waiting interrupt signal.")
	<-doneChan
}
