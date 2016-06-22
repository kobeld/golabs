package main

import (
	"strconv"

	"fmt"
)

func main() {
	floA := 35.69146649
	floB := 33.867487

	fmt.Printf("%.6f\n", floA)
	fmt.Printf("%.6f\n", floB)

	floStr := fmt.Sprintf("%.6f", floA+floB)
	floStrLen := len(floStr)
	println(floStr)
	println(floStr[floStrLen-3 : floStrLen])

	code, _ := strconv.Atoi(floStr[floStrLen-3 : floStrLen])
	println(code)
}
