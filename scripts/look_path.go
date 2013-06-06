package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	appName := "ls"
	path, err := exec.LookPath(appName)
	if err != nil {
		log.Fatal("Cound not find path of %s", appName)
	}
	fmt.Printf("'%s' is available at %s\n", appName, path)
}
