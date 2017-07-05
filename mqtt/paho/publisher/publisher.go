package main

import (
	"time"

	mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func main() {
	var (
		// serverAddr = "tcp://api.iguiyu.com:1883"
		// id          = "560a5c34b51b8e0001000001"
		serverAddr  = "tcp://192.168.99.100:1883"
		id          = "560a00fc63ed2e4f8c000004"
		topicRoot   = "hud/" + id
		topicReport = topicRoot + "/_reserved/report"
		topicStatus = topicRoot + "/_reserved/status"
		opts        = mqtt.NewClientOptions().AddBroker(serverAddr).SetClientID(id).SetCleanSession(true).SetWill(topicStatus, "off", 0, true)
	)

	opts = opts.SetAutoReconnect(true)

	// opts.SetConnectionLostHandler(func(c *mqtt.Client, err error) {
	// 	println("------ Is disconnected! ------")
	// 	time.Sleep(20 * time.Second)
	// 	println("------ Time -----")
	// 	c.Connect()
	// })

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Send online status
	if token := client.Publish(topicStatus, 0, true, "on"); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	report := []byte(`
		{"deviceId": "560a00fc63ed2e4f8c000004", "rpm": "3100", "spd": "120", "fli": "69", "xm":"8", "reportedAt": "1445485125599"}
		`)

	if token := client.Publish(topicReport, 1, false, report); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	time.Sleep(3 * time.Second)

	for client.IsConnected() {
	}

	println("Stop")

	// if token := client.Publish(topicStatus, 0, true, "off"); token.Wait() && token.Error() != nil {
	// 	panic(token.Error())
	// }

	// time.Sleep(1 * time.Second)

	// Disconnect gracefully
	// client.Unsubscribe(topicReport, topicStatus)
	// client.Disconnect(250)
}
