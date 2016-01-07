package main

import (
	"time"

	"github.com/kobeld/goutils"

	"git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

type MqttDomain interface {
	Address() string
	ClientId() string
	Topics() []TopicItem
}

type TopicItem struct {
	Topic    string
	Qos      byte
	Callback func(msg mqtt.Message)
}

type MessageHandler func(msg mqtt.Message)

type MqttManager struct {
	Domain MqttDomain
	Client *mqtt.Client
}

func (this *MqttManager) Run() (err error) {
	if this.Client == nil {
		return
	}

	token := this.Client.Connect()
	if token.Wait() && token.Error() != nil {
		err = token.Error()
		return
	}

	return
}

func NewMqttManager(domain MqttDomain) (manager *MqttManager, err error) {

	opts := &mqtt.ClientOptions{
		ClientID:             domain.ClientId(),
		CleanSession:         true,
		MaxReconnectInterval: 1 * time.Second,
		AutoReconnect:        true,
		KeepAlive:            30 * time.Second,
	}

	opts.AddBroker(domain.Address())
	opts.OnConnect = func(c *mqtt.Client) {
		println("Connected")

		for _, item := range manager.Domain.Topics() {

			println("Reregistering topic: " + item.Topic)

			// TODO: Should also store the qos
			token := manager.Client.Subscribe(item.Topic, item.Qos, func(c *mqtt.Client, msg mqtt.Message) {
				item.Callback(msg)
			})
			if token.Wait() && token.Error() != nil {
				goutils.PrintStackAndError(token.Error())
				continue
			}

		}
	}

	manager = &MqttManager{
		Domain: domain,
		Client: mqtt.NewClient(opts),
	}

	return
}
