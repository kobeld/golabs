package main

import (
	"encoding/json"
	"github.com/bitly/nsq/nsq"
	"github.com/theplant/qortex/nsqproducers"
	"log"
	"time"
)

type LogAsyncHandler struct{}

func (this *LogAsyncHandler) HandleMessage(msg *nsq.Message, responseChan chan *nsq.FinishedMessage) {
	time.Sleep(time.Second * 5)
	log.Printf("-------- Receiving new message: %s \n", msg.Body)
	<-responseChan
	return
}

type LogHandler struct{}

func (this *LogHandler) HandleMessage(msg *nsq.Message) error {
	time.Sleep(time.Second * 1)

	entryData := new(nsqproducers.EntryData)
	err := json.Unmarshal(msg.Body, &entryData)
	if err != nil {
		panic(err)
	}

	log.Printf("-------- Receiving new message: %+v\n", entryData)
	return nil
}

func main() {

	topic := "create_entry"
	channel := "testing"
	lookupAddr := "localhost:4161"

	reader, err := nsq.NewReader(topic, channel)
	if err != nil {
		panic(err)
	}

	// reader.AddAsyncHandler(&LogHandler{})

	reader.AddHandler(&LogHandler{})

	err = reader.ConnectToLookupd(lookupAddr)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-reader.ExitChan:
			return
		}
	}

}
