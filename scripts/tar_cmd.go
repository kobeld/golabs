package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	inputReader *bufio.Reader
	input       string
	err         error
)

func main() {
	cmd := exec.Command("tr", "a-z", "A-Z")

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input:")

	input, err = inputReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	cmd.Stdin = strings.NewReader(input)

	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The reslut of all caps: %q\n", out.String())
}
