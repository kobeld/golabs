package main

var i = "Vars come first"

func init() {
	println(i)
	println("Second is init func")
}

func main() {
	println("Now should be the execution")
}
