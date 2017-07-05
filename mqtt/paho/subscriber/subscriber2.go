package main

import (
	"fmt"
	"time"

	mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	var (
		serverAddr = "tcp://192.168.99.100:1883"
		id         = "Hud Server " + bson.NewObjectId().Hex()

		topic  = "wave/ent100"
		topic2 = "wave/pleth"

		opts = mqtt.NewClientOptions().AddBroker(serverAddr).SetClientID(id).SetCleanSession(true)
	)

	client := mqtt.NewClient(opts)

	go func() {

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		if token := client.Subscribe(topic, 1, topic1Message); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		if token := client.Subscribe(topic2, 1, topic2Message); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

	}()

	for {
	}

	client.Unsubscribe(topic)
	client.Disconnect(250)
}

func topic1Message(client *mqtt.Client, msg mqtt.Message) {
	println("----- Status -----")
	println(string(msg.Payload()))
}

func topic2Message(client *mqtt.Client, msg mqtt.Message) {
	println("----- Status -----")
	println(string(msg.Payload()))
}

func reportMessage(client *mqtt.Client, msg mqtt.Message) {

	go func() {
		println("----- Report -----")
		fmt.Printf("%+v\n", msg.Topic())
		println(string(msg.Payload()))
		println(time.Now().String())
		time.Sleep(5 * time.Second)
	}()
}
