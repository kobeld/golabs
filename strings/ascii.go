package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "Hello world"
	a_asc := strconv.QuoteToASCII(a)

	b := "你好"
	b_asc := strconv.QuoteToASCII(b)

	println(len(a))
	println(len(a_asc))
	println(a)
	println(a_asc)

	println(a[0])

	println(len(b))
	println(len(b_asc))

	println(b)
	println(b_asc)

	s := strconv.QuoteToASCII(`"Fran & Freddie's Diner	☺"`)
	println(s)
	fmt.Println(s)
}
