package ExternalDevice

import (
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func GetTopicPrefix(d Device) string {
	uID := d.GetUniqueId()
	if uID != "" {
		uID = uID + "/"
	}
	return NodeID + "/" + d.GetRawId() + "/" + uID
}

func GetDiscoveryTopic(d Device) string {
	uID := d.GetUniqueId()
	if uID != "" {
		uID = uID + "/"
	}
	return DiscoveryPrefix + "/" + d.GetRawId() + "/" + NodeID + "/" + uID
}

func GetTopic(d Device, rawTopicString string) string {
	return GetTopicPrefix(d) + strings.TrimSuffix(rawTopicString, "_topic")
}

func MakeMessageHandler(d Device) func(client mqtt.Client, msg mqtt.Message) {

	return func(client mqtt.Client, msg mqtt.Message) {

		topicFound := false

		for topic, f := range topicStore {
			if msg.Topic() == topic {
				topicFound = true
				(*f)(msg, client)
				d.UpdateState()
			}
		}

		if !topicFound {
			log.Println("Unknown Message on topic " + msg.Topic())
			log.Println(msg.Payload())
		}
	}

}

type MQTTFields struct {
	Client         *mqtt.Client
	ForceUpdate    bool
	MessageHandler mqtt.MessageHandler
	UpdateInterval float64
}
