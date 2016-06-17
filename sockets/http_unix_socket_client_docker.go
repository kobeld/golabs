package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tv42/httpunix"
)

var (
	sockFile = "/var/run/docker.sock"
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

	resp, err := client.Get("http+unix://myservice/info")
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	bodystr := string(body)
	fmt.Println(bodystr)

}
