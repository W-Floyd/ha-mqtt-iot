package hadiscovery

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//go:generate go run ../helpers/

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

type store struct {
	BinarySensor map[string]string
	Light        struct {
		BrightnessState map[string]string
		ColorTempState  map[string]string
		EffectState     map[string]string
		HsState         map[string]string
		RgbState        map[string]string
		State           map[string]string
		WhiteValueState map[string]string
		XyState         map[string]string
	}
	Sensor map[string]string
	Switch map[string]string
}

func initStore() store {
	var s store
	s.BinarySensor = make(map[string]string)
	s.Light.BrightnessState = make(map[string]string)
	s.Light.ColorTempState = make(map[string]string)
	s.Light.EffectState = make(map[string]string)
	s.Light.HsState = make(map[string]string)
	s.Light.RgbState = make(map[string]string)
	s.Light.State = make(map[string]string)
	s.Light.WhiteValueState = make(map[string]string)
	s.Light.XyState = make(map[string]string)
	s.Sensor = make(map[string]string)
	s.Switch = make(map[string]string)
	return s
}

var stateStore = initStore()
