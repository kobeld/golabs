package main

import (
	"fmt"
	"time"
)

func main() {

	// Can uncomment this to get a serial result.
	// runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Print(" - ")
		}

	}()

	for i := 0; i < 1000; i++ {
		fmt.Print("*")
	}

	time.Sleep(1e9)
}
