package config

import (
	"log"
	"time"

	ExternalDevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (c Config) Convert() ([]ExternalDevice.Device, *mqtt.ClientOptions) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(c.MQTT.Broker)
	opts.SetUsername(c.MQTT.Username)
	opts.SetPassword(c.MQTT.Password)
	opts.SetKeepAlive(30 * time.Second)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("TOPIC: %s\n", msg.Topic())
		log.Printf("MSG: %s\n", msg.Payload())
	})
	opts.SetPingTimeout(1 * time.Second)
	opts.SetAutoReconnect(true)
	if c.MQTT.NodeID != "" {
		ExternalDevice.NodeID = c.MQTT.NodeID
	}
	opts.SetClientID(ExternalDevice.NodeID)
	if c.MQTT.InstanceName != "" {
		ExternalDevice.InstanceName = c.MQTT.InstanceName
	}

	return c.Translate(), opts
}
