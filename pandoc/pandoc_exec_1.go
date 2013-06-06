package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {

	htmlDoc := "<strong>Hello World!</strong>"

	cmd := exec.Command("pandoc", "-f", "html", "-t", "markdown")
	cmd.Stdin = strings.NewReader(htmlDoc)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The result is: %q \n", out.String())
}
