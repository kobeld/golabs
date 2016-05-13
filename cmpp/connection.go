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
	connectTimeout time.Duration = time.Second * 2
)

func main() {

	client := cmpp.NewClient(cmpp.V30)
	defer client.Disconnect()

	// err := client.Connect("183.230.96.94:17890", user, pwd, connectTimeout)
	err := client.Connect("123.59.129.126:8080", user, pwd, connectTimeout)
	// err := client.Connect("localhost:5000", user, pwd, connectTimeout)
	if goutils.HasErrorAndPrintStack(err) {
		return
	}
	log.Println("Client connect and auth ok!")

	for {
		i, err := client.RecvAndUnpackPkt(0)
		if goutils.HasErrorAndPrintStack(err) {
			break
		}

		fmt.Printf("%+v\n", i)
	}

	// time.Sleep(5 * time.Second)

	// for {

	// 	i, err := client.RecvAndUnpackPkt(0)
	// 	if goutils.HasErrorAndPrintStack(err) {
	// 		break
	// 	}

	// 	switch p := i.(type) {
	// 	case *cmpp.Cmpp3SubmitRspPkt:
	// 		log.Printf("receive a cmpp3 submit response: %v.", p)

	// 	case *cmpp.CmppActiveTestReqPkt:
	// 		log.Printf("client %d: receive a cmpp active request: %v.", p)
	// 		// rsp := &cmpp.CmppActiveTestRspPkt{}
	// 		// err := client.SendRspPkt(rsp, p.SeqId)
	// 		// if goutils.HasErrorAndPrintStack(err) {
	// 		// 	break
	// 		// }

	// 	case *cmpp.CmppActiveTestRspPkt:
	// 		log.Printf("receive a cmpp activetest response: %v.", p)

	// 	case *cmpp.CmppTerminateReqPkt:
	// 		log.Printf("receive a cmpp terminate request: %v.", p)

	// 	case *cmpp.CmppTerminateRspPkt:
	// 		log.Printf("client %d: receive a cmpp terminate response: %v.", p)
	// 	}

	// }

	log.Println("End!")
}
