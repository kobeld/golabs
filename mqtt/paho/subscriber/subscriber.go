package main

import (
	"fmt"
	"time"

	mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	var (
		serverAddr  = "tcp://192.168.99.100:1883"
		id          = "Hud Server " + bson.NewObjectId().Hex()
		topicRoot   = "/hud/+"
		topicReport = topicRoot + "/report"
		topicStatus = topicRoot + "/status"
		opts        = mqtt.NewClientOptions().AddBroker(serverAddr).SetClientID(id).SetCleanSession(true)
	)

	client := mqtt.NewClient(opts)

	go func() {

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		if token := client.Subscribe(topicStatus, 1, statusMessage); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		if token := client.Subscribe(topicReport, 1, reportMessage); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

	}()

	for {
	}

	client.Unsubscribe(topicStatus, topicReport)
	client.Disconnect(250)
}

func statusMessage(client *mqtt.Client, msg mqtt.Message) {
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
