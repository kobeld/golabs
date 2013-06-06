package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ssh", "-v", "ubuntu@qortex.theplant-dev.com", "mkdir", "makefromgo")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	println("Done!")
}
