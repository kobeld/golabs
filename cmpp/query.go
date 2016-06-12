package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kobeld/goutils"

	cmpp "github.com/bigwhite/gocmpp"
)

const (
	user           string        = "150026"
	pwd            string        = "150026"
	connectTimeout time.Duration = time.Second * 20
)

func main() {

	client := cmpp.NewClient(cmpp.V30)
	defer client.Disconnect()

	err := client.Connect("123.59.129.126:8080", user, pwd, connectTimeout)
	// err := client.Connect("localhost:5000", user, pwd, connectTimeout)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}
	log.Println("Client connect and auth ok!")

	go func() {
		time.Sleep(3 * time.Second)

		println("hehehehehe")
		p := &cmpp.CmppQueryReqPkt{
			QueryAt: "20160605",
			Type:    0,
			// Code:    "BJYG",
		}

		println("Start to send req")
		err = client.SendReqPkt(p)
		if err != nil {
			log.Printf("send a cmpp3 query request error: %s.", err)
		} else {
			log.Printf("send a cmpp3 query request ok")
		}
	}()

	for {

		println("Receiving.....")
		i, err := client.RecvAndUnpackPkt(0)
		if goutils.HasErrorAndPrintStack(err) {
			break
		}

		fmt.Printf("%+v\n", i)
	}

	log.Println("End!")
}
