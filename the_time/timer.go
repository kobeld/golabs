package main

import "time"

func main() {

	timer := time.NewTimer(5 * time.Second)

	for _ = range timer.C {
		println("hello")
		timer.Reset(5 * time.Second)
	}

}
