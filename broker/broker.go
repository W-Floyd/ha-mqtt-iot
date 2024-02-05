package broker

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Broker interface {
	Connect()
	Disconnect()
	Messages() chan mqtt.Message
}
