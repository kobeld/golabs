package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "Hello world!")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
