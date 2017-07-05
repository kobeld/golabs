package main

import (
	"time"

	mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func main() {
	var (
		serverAddr = "tcp://192.168.99.100:1883"
		id         = "560a00fc63ed2e4f8c000004"
		topic1     = "wave/invp1"
		topic2     = "wave/ent100"
		topic3     = "wave/ecg1"
		topic4     = "wave/pleth"
		opts       = mqtt.NewClientOptions().AddBroker(serverAddr).SetClientID(id).SetCleanSession(true)
	)

	opts = opts.SetAutoReconnect(true)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Publish(topic1, 0, false, "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Publish(topic2, 0, false, "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Publish(topic3, 0, false, "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Publish(topic4, 0, false, "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	time.Sleep(1 * time.Second)

	println("Stop")

}
