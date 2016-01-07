package main

import (
	"fmt"
	"log"

	"git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

type HudDomain struct {
}

func (this *HudDomain) Address() string {
	return "tcp://192.168.99.100:1883"
}

func (this *HudDomain) ClientId() string {
	return "Hud server"
}

func (this *HudDomain) Topics() (items []TopicItem) {
	// r := make(map[string]MessageHandler)

	// r["hud/+/_reserved/status"] = func(msg mqtt.Message) {
	// 	log.Println("----- Receive Hud Report Message-----")
	// 	fmt.Printf("%+v\n", msg)
	// }

	item := TopicItem{
		Topic: "hud/+/_reserved/status",
		Qos:   0,
		Callback: func(msg mqtt.Message) {
			log.Println("----- Receive Hud Report Message-----")
			fmt.Printf("%+v\n", msg)
		},
	}

	items = append(items, item)

	return
}

func main() {

	var (
		domain = new(HudDomain)
		// manager, err =
	)

	manager, err := NewMqttManager(domain)
	if err != nil {
		panic(err)
	}

	manager.Run()

	for {
	}
}
