package main

import (
	"time"
)

func main() {
	ch := make(chan int)
	go pump(ch)
	go suck(ch)
	time.Sleep(1e9)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		println(<-ch)
	}
}
