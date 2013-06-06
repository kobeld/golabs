package main

import (
	"fmt"
	"github.com/nutrun/lentil"
)

const (
	addr = "127.0.0.1:11300"
)

func main() {
	conn, err := lentil.Dial(addr)
	if err != nil {
		panic(err)
	}

	tubes, err := conn.ListTubes()
	if err != nil {
		panic(err)
	}

	for _, tube := range tubes {
		println(tube)
		m, _ := conn.StatsTube(tube)
		fmt.Printf("%+v \n", m)
	}
}
