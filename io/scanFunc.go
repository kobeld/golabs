package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Comma-separated list; last entry is empty.
	// const input = "{'name':'Aaron'}{'name':'Jane'}"

	file, err := os.Open("file.data")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// scanner := bufio.NewScanner(strings.NewReader(input))
	// Define a split function that separates on commas.
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// for i := 0; i < len(data); i++ {
		// 	if data[i] == ',' {
		// 		return i + 1, data[:i], nil
		// 	}
		// }

		size := len(data)

		for i := 0; i < size; i++ {

			if i+1 < size {
				if data[i] == '}' && data[i+1] == '{' {
					return i + 1, data[:i+1], nil
				}

			}

		}

		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	// Scan.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
