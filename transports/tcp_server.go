package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/parnurzeal/gorequest"

	"github.com/kobeld/goutils"

	"net"
)

var (
	activatedMap = make(map[string]string)

	mobile = flag.String("mobile", `1064848022012`, "the mobile")
)

func main() {

	flag.Parse()

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	println("Message will be sent to: " + *mobile)
	println("Listen and serve on :8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		println("Connection Accepted!")

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	for {

		buf := make([]byte, 128)

		n, err := c.Read(buf)
		if goutils.HasErrorAndPrintStack(err) {
			return
		}

		fmt.Printf("Read %d\n", n)

		if n == 11 {
			var ping = string(buf[0:11])

			fmt.Printf("Received message: %s\n", ping)

			if ping == "*Login#1603" {

				log.Println("Login and Reset map.")
				activatedMap = make(map[string]string)
			}
			continue
		}

		var (
			toMsg string
			msg   = string(buf[0:17])
			info  = strings.Split(msg, ",")
		)

		fmt.Printf("Received message: %s\n", msg)

		if len(info) < 3 {
			log.Println("Ignore message and continue")
			continue
		}

		if info[1] != "$053" {
			continue
		}

		var (
			// status string
			sId = info[3]
			// last   = info[len(info)-1]
			status = info[len(info)-1]
		)

		// fmt.Printf("The last is %s, and the lenth is %d\n", last, len(last))

		// if len(last) > 0 {
		// 	status = string(last[0])
		// }

		if status != "0" && status != "1" {

			// fmt.Printf("The status is %s.\n", status)
			fmt.Printf("The status is %s, and the lenth is %d\n", status, len(status))
			println("Unknow status, ignored!")
			continue
		}

		toMsg = fmt.Sprintf("#000,#r%s", sId)

		oldStatus, ok := activatedMap[sId]
		if !ok {
			println("Send message: " + toMsg)
			_, err = fmt.Fprintf(c, toMsg)
			if goutils.HasErrorAndPrintStack(err) {
				return
			}
		}

		activatedMap[sId] = status

		if oldStatus != status && sId == "58" {
			// Make a signal
			url := "http://parking.iguiyu.com/sms-test"

			goutils.CoveredGo(func() {
				_, _, errs := gorequest.New().Post(url).
					Send("mobile=" + *mobile).
					Send("content=" + fmt.Sprintf("S%s", status)).
					End()

				for _, err := range errs {
					println("Err:" + err.Error())
				}
			})
		}
	}
}
