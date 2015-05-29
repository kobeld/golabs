package main

import (
	"fmt"
)

func main() {

	arr := []string{"one", "", "two", "three", "", "four", "two", ""}

	for i, ele := range arr {
		if ele == "two" {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}

	fmt.Printf("%+v\n", arr)
}
