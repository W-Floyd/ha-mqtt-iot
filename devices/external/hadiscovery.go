package ExternalDevice

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	retain bool = true
	qos    byte = 0
)

var topicStore = make(map[string]*func(mqtt.Message, mqtt.Client))

// DiscoveryPrefix is the prefix that HA uses to discover on.
// Does not usually need changing.
const DiscoveryPrefix = "homeassistant"

// SWVersion is the software version.
// TODO - Move this elsewhere maybe?
const SWVersion = "0.4.5"

var SoftwareName = "Homeassistant MQTT IOT"

// InstanceName is the instance name, helpful for identifying a given client
var InstanceName = "Homeassistant MQTT IOT"

// NodeID is the Node ID, that is, what that node connects under.
var NodeID = "ha-mqtt-iot"

var Manufacturer = "William Floyd"

///////////////////

var stateStore = initStore()
