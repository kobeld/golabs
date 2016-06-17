package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tv42/httpunix"
)

var (
	sockFile = os.Getenv("GOPATH") + "/src/github.com/kobeld/golabs/sockets/http.sock"
)

func main() {

	u := &httpunix.Transport{
		DialTimeout:           100 * time.Millisecond,
		RequestTimeout:        1 * time.Second,
		ResponseHeaderTimeout: 1 * time.Second,
	}
	u.RegisterLocation("myservice", sockFile)

	var client = http.Client{
		Transport: u,
	}

	resp, err := client.Get("http+unix://myservice/hello")
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	bodystr := string(body)
	fmt.Println(bodystr)

	// buf, err := httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// println(len(buf))

	// fmt.Printf("%s", buf)
	// resp.Body.Close()

}
