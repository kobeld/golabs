package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)

	// var jsonBlob = []byte(`{"Name": "Quoll",    "Order": "Dasyuromorphia"}`)
	type Animal struct {
		Name     string
		Order    string
		Location string
	}

	animal := Animal{
		Location: "Hangzhou",
	}

	animals := []Animal{animal, Animal{}}

	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}
