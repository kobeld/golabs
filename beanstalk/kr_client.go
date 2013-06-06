package main

import (
	"github.com/kr/beanstalk"
	"time"
)

func main() {
	conn, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		panic(err)
	}

	id, err := conn.Put([]byte("hello"), 1, 0, 120*time.Second)
	if err != nil {
		panic(err)
	}

	println("Job ID: ", id)

	for {
		rid, body, err := conn.Reserve(5 * time.Second)
		if err != nil {
			println(err)
		}

		println("Reserved job ID: ", rid)
		println(string(body))
	}

}
