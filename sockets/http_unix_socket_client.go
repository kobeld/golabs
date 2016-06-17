package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"net"
)

var (
	sockFile = os.Getenv("GOPATH") + "/src/github.com/kobeld/golabs/sockets/http.sock"
)

func fakeDial(proto, addr string) (conn net.Conn, err error) {

	return net.Dial("unix", sockFile)

}

func main() {

	http.Get(url)

	tr := &http.Transport{
		Dial: fakeDial,
	}

	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get("http://any/hello")
	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	bodystr := string(body)
	fmt.Println(bodystr)
}
