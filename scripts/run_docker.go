package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("docker", "run", "--rm", "-v",
		"/Users/kobeld/mygo/src/github.com/iguiyu/neuro/web/public/files:/tmp", "kobeld/wkhtmltopdf", "http://192.168.199.109:9000/orders/56d57e9963ed2e6d9722ea4f/pdf", "/tmp/signed.pdf")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	println("Done!")
}

// docker run --rm -v
// /Users/kobeld/Desktop:/tmp kobeld/wkhtmltopdf http://192.168.10.108:9000/orders/56d57e9963ed2e6d9722ea4f/preview /tmp/signature.pdf
