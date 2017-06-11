package main

func main() {
	arr := []byte{1, 2, 3, 4, 5}

	println(arr)

	for index, _ := range arr {
		println(&arr[index])
	}

	println()

	arr2 := []string{"hello", "aaron", "there"}

	for index, _ := range arr2 {
		println(&arr2[index])
	}

	map1 := map[string]string{"1": "a", "2": "b", "3": "c"}
	for key, _ := range map1 {
		println(&map1[key])
	}

}
