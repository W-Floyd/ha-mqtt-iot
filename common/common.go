package common

func StringPointer(input string) *string {
	return &input
}

// DiscoveryPrefix is the prefix that HA uses to discover on.
// Does not usually need changing.
const DiscoveryPrefix = "homeassistant"

// SWVersion is the software version.
// TODO - Move this elsewhere maybe?
const SWVersion = "0.4.4"

// InstanceName is the instance name, helpful for identifying a given client
var InstanceName = "Homeassistant MQTT IOT"

// NodeID is the Node ID, that is, what that node connects under.
var NodeID = "ha-mqtt-iot"
