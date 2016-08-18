package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Println(os.Getpid())
	fmt.Println(os.Getppid())
	syscall.Getpid()
}
