package InternalDevice

import externaldevice "github.com/W-Floyd/ha-mqtt-iot/devices/externaldevice"

////////////////////////////////////////////////////////////////////////////////
// Do not modify this file, it is automatically generated
////////////////////////////////////////////////////////////////////////////////
//
type DeviceTracker struct {
	Devices        *([]string) `json:"devices,omitempty"`          // "List of devices with their topic."
	PayloadHome    *string     `json:"payload_home,omitempty"`     // "The payload value that represents the 'home' state for the device."
	PayloadNotHome *string     `json:"payload_not_home,omitempty"` // "The payload value that represents the 'not_home' state for the device."
	Qos            *int        `json:"qos,omitempty"`              // "The QoS level of the topic."
	SourceType     *string     `json:"source_type,omitempty"`      // "Attribute of a device tracker that affects state when being used to track a [person](/integrations/person/). Valid options are `gps`, `router`, `bluetooth`, or `bluetooth_le`."
	MQTT           struct {
		UpdateInterval *float64 `json:"update_interval,omitempty"`
		ForceUpdate    *bool    `json:"force_update,omitempty"`
	} `json:"mqtt"`
}

func (iDevice DeviceTracker) Translate() externaldevice.DeviceTracker {
	eDevice := externaldevice.DeviceTracker{}
	eDevice.MQTT = new(externaldevice.MQTTFields)
	if iDevice.MQTT.ForceUpdate != nil {
		eDevice.MQTT.ForceUpdate = iDevice.MQTT.ForceUpdate
	}
	if iDevice.MQTT.UpdateInterval != nil {
		eDevice.MQTT.UpdateInterval = iDevice.MQTT.UpdateInterval
	}
	if iDevice.Devices != nil {
		eDevice.Devices = iDevice.Devices
	}
	if iDevice.PayloadHome != nil {
		eDevice.PayloadHome = iDevice.PayloadHome
	}
	if iDevice.PayloadNotHome != nil {
		eDevice.PayloadNotHome = iDevice.PayloadNotHome
	}
	if iDevice.Qos != nil {
		eDevice.Qos = iDevice.Qos
	}
	if iDevice.SourceType != nil {
		eDevice.SourceType = iDevice.SourceType
	}
	eDevice.Initialize()
	return eDevice
}
