package main

func main() {
	var (
		total   = 10
		part    = 9
		compare = 80
	)

	result := float32(part) / float32(total) * 100
	println(result)

	if float32(compare) >= result {
		println("true")
	} else {
		println("false")
	}
}
