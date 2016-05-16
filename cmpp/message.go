package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kobeld/goutils"

	cmpp "github.com/bigwhite/gocmpp"
	"github.com/bigwhite/gocmpp/utils"
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
		time.Sleep(2 * time.Second)

		println("hehehehehe")
		//submit a message
		cont, err := cmpputils.Utf8ToUcs2("gocmpp submit")
		if err != nil {
			fmt.Printf("utf8 to ucs2 transform err: %s.", err)
			return
		}
		p := &cmpp.Cmpp3SubmitReqPkt{
			PkTotal:            1,
			PkNumber:           1,
			RegisteredDelivery: 1,
			MsgLevel:           1,
			ServiceId:          "BJYG",
			FeeUserType:        2,
			// FeeTerminalId:      "",
			FeeTerminalType:  1,
			MsgFmt:           8,
			MsgSrc:           "150026",
			FeeType:          "01",
			FeeCode:          "11",
			ValidTime:        "",
			AtTime:           "",
			SrcId:            "1064899150026",
			DestUsrTl:        1,
			DestTerminalId:   []string{"1064820137721"},
			DestTerminalType: 1,
			MsgLength:        uint8(len(cont)),
			MsgContent:       cont,
		}

		println("Start to send req")
		err = client.SendReqPkt(p)
		if err != nil {
			log.Printf("send a cmpp3 submit request error: %s.", err)
		} else {
			log.Printf("send a cmpp3 submit request ok")
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
