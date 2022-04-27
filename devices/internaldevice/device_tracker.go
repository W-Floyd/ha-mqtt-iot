package InternalDevice

import externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
func (iDevice DeviceTracker) Translate() externaldevice.DeviceTracker {
	eDevice := externaldevice.DeviceTracker{}
	eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	eDevice.Devices = iDevice.Devices
	eDevice.PayloadHome = iDevice.PayloadHome
	eDevice.PayloadNotHome = iDevice.PayloadNotHome
	eDevice.Qos = iDevice.Qos
	eDevice.SourceType = iDevice.SourceType
	eDevice.Initialize()
	return eDevice
}

type DeviceTracker struct {
	Devices        []string `json:"devices"`          // "List of devices with their topic."
	PayloadHome    string   `json:"payload_home"`     // "The payload value that represents the 'home' state for the device."
	PayloadNotHome string   `json:"payload_not_home"` // "The payload value that represents the 'not_home' state for the device."
	Qos            int      `json:"qos"`              // "The QoS level of the topic."
	SourceType     string   `json:"source_type"`      // "Attribute of a device tracker that affects state when being used to track a [person](/integrations/person/). Valid options are `gps`, `router`, `bluetooth`, or `bluetooth_le`."
	MQTT           struct {
		UpdateInterval float64 `json:"update_interval"`
		ForceUpdate    bool    `json:"force_update"`
	} `json:"mqtt"`
}
