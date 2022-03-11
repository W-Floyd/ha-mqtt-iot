package hadiscovery

import mqtt "github.com/eclipse/paho.mqtt.golang"

type MQTTFields struct {
	Client         *mqtt.Client
	MessageHandler mqtt.MessageHandler
	UpdateInterval float64
	ForceUpdate    bool
}
