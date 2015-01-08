package main

import (
	"log"
	"time"
)

func main() {

	ticker := time.NewTicker(1 * time.Second)

	go run(ticker)

	n := 0
	for {
		time.Sleep(2 * time.Second)
		n += 1
		println(n)

		if n == 3 {
			ticker.Stop()
		}
	}

}

func run(ticker *time.Ticker) {
	log.Println("hello")
	for {
		println("------")
		select {
		case data, ok := <-ticker.C:

			log.Printf("%+v \n", data)

			if ok {
				println("I am ok")
			}

			// simply push blank message
			println("inside")
		}
		println("******")
	}

	println("&&&&&&&&&")

	return
}
