package iotconfig

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/W-Floyd/ha-mqtt-iot/iotconfig/config"
)

type Config config.Config

// Convert takes a config and turns it into a format that can be used for HA.
func (sconfig Config) Convert() (opts *mqtt.ClientOptions, lights []ExternalDevice.Light) {
	opts = mqtt.NewClientOptions()
	opts.AddBroker(sconfig.MQTT.Broker)
	opts.SetUsername(sconfig.MQTT.Username)
	opts.SetPassword(sconfig.MQTT.Password)
	opts.SetKeepAlive(30 * time.Second)
	opts.SetDefaultPublishHandler(func(_ mqtt.Client, msg mqtt.Message) {
		log.Printf("TOPIC: %s\n", msg.Topic())
		log.Printf("MSG: %s\n", msg.Payload())
	})
	opts.SetPingTimeout(1 * time.Second)
	opts.SetAutoReconnect(true)
	if sconfig.MQTT.NodeID != "" {
		hadiscovery.NodeID = sconfig.MQTT.NodeID
	}
	opts.SetClientID(hadiscovery.NodeID)
	if sconfig.MQTT.InstanceName != "" {
		hadiscovery.InstanceName = sconfig.MQTT.InstanceName
	}

	for _, li := range sconfig.Lights {
		lights = append(lights, li.Translate())
	}

	return
}
