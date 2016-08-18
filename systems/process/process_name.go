package main

import (
	"fmt"
	"os"
)

// 注：编译后执行的结果跟直接 go run 的结果有差别
func main() {
	processName := os.Args[0]
	fmt.Println(processName)
}
