package main

import (
	"fmt"
	"github.com/iwanbk/gobeanstalk"
)

const (
	addr = "127.0.0.1:11300"
)

func main() {
	conn, err := gobeanstalk.Dial(addr)
	if err != nil {
		panic(err)
	}

	// id, err := conn.Put([]byte("Hello Beanstalk"), 0, 0, 10)
	// if err != nil {
	// 	panic(err)
	// }

	// println("Job is inserted, ID: ", id)

	for {
		j, err := conn.Reserve()
		if err != nil {
			panic(err)
		}

		fmt.Printf("id:%d, body:%s\n", j.Id, string(j.Body))

		err = conn.Delete(j.Id)
		if err != nil {
			panic(err)
		}
	}

}
