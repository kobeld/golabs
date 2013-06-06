package main

import (
	"time"
)

func main() {
	ch := make(chan int)
	go send(ch)
	go handle(ch)
	time.Sleep(1 * 1e7)
	close(ch)
	time.Sleep(1 * 1e7)
	println("end main")
}

func send(ch chan<- int) {
	for {
		ch <- 1
	}

	println("End send")
}

func handle(ch <-chan int) {
	for d := range ch {
		print(d)
	}

	println("End Handle")
}

func next() chan<- int {
	ch := make(chan<- int)

	return ch

}
