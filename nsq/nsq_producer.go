package main

import (
	"github.com/bitly/nsq/nsq"
	"log"
)

func main() {
	topic := "setup_org"
	lookupAddr := "localhost:4150"

	writer := nsq.NewWriter(lookupAddr)

	orgs := []string{"The Plant", "iA", "Duoerl", "Kobeld"}

	for _, org := range orgs {
		frameType, data, err := writer.Publish(topic, []byte(org))

		log.Printf("FrameType: %+v . Data: %+v \n", frameType, string(data))
		if err != nil {
			println(err.Error())
		}

	}

}
