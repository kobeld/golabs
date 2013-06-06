package main

import (
	"github.com/kobeld/xmpp"
	"os"
)

func main() {
	xmppConfig := &xmpp.Config{
		Log:   os.Stdout,
		InLog: os.Stdout,
		// OutLog:         os.Stdout,
		Create:         false,
		Archive:        false,
		TrustedAddress: true,
	}

	address := "qortex.theplant-dev.com:5222"
	user := "kobeld"
	domain := "qortex.theplant-dev.com"
	abc := "bleach"

	stanza := make(chan xmpp.Stanza)

	conn, err := xmpp.Dial(address, user, domain, abc, xmppConfig)
	if err != nil {
		panic(err)
	}

	rosterReply, _, err := conn.RequestRoster()
	if err != nil {
		panic(err)
	}

	conn.SignalPresence("")

	go func() {
		for {
			stan, err := conn.Next()
			if err != nil {
				panic(err)
				continue
			}

			stanza <- stan
		}
	}()

	for {
		select {
		case s := <-stanza:
			println("------")
			println(s.Name.Local)
		case roster := <-rosterReply:
			rosterEntries, _ := xmpp.ParseRoster(roster)
			for _, re := range rosterEntries {
				println(re.Jid)
			}
		}
	}
}
