package main

import (
	"time"
)

func main() {
	current := time.Now()
	println(current.String())
	println(current.Truncate(24 * time.Hour).UTC().String())
}
