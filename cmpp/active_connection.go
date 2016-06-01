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
	if goutils.HasErrorAndPrintStack(err) {
		return
	}

	log.Println("Client connect and auth ok!")

	var (
		pingTimes = 0
		data      = make(chan interface{}, 1)
		timer     = time.NewTimer(3 * time.Minute)
	)

	go func() {
		for {
			i, err := client.RecvAndUnpackPkt(0)
			if goutils.HasErrorAndPrintStack(err) {
				break
			}

			data <- i
		}
	}()

	for {
		select {
		case i := <-data:
			fmt.Printf("===== Incomming data: \n%+v\n=====\n\n", i)

			timer.Reset(3 * time.Minute)
			pingTimes = 0

			switch p := i.(type) {
			case *cmpp.Cmpp3DeliverReqPkt:
				log.Printf("Received a cmpp3 deliver request:\n %+v.\n", p)

				rsp := &cmpp.Cmpp3DeliverRspPkt{
					MsgId:  p.MsgId,
					Result: 0,
				}

				err := client.SendRspPkt(rsp, p.SeqId)
				if goutils.HasErrorAndPrintStack(err) {
					continue
				}

			case *cmpp.Cmpp3SubmitRspPkt:
				log.Printf("Received a submit response:\n %+v.\n", p)

			case *cmpp.CmppActiveTestReqPkt:
				log.Printf("Received an active request:\n %+v.\n", p)
				rsp := &cmpp.CmppActiveTestRspPkt{}
				err := client.SendRspPkt(rsp, p.SeqId)
				if goutils.HasErrorAndPrintStack(err) {
					continue
				}

			case *cmpp.CmppActiveTestRspPkt:
				log.Printf("Received an active response:\n %+v.\n", p)

			case *cmpp.CmppTerminateReqPkt:
				log.Printf("Received a terminate request:\n %+v.\n", p)

			case *cmpp.CmppTerminateRspPkt:
				log.Printf("Received a terminate response:\n %+v.\n", p)
			}

		case <-timer.C:
			// TODO:
			println("No data in 3 munites, sending active message to gateway")
			err := client.SendReqPkt(new(cmpp.CmppActiveTestReqPkt))
			if goutils.HasErrorAndPrintStack(err) {
				continue
			}

			pingTimes++
			if pingTimes < 3 {
				timer.Reset(1 * time.Minute)
			}
		}
	}

	log.Println("End!")
}
